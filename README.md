# CLO Go client

The official Go client for the [CLO](https://clo.ru) API.

As of **v3**, the client is **generated from the CLO OpenAPI spec** ([api.clo.ru/openapi.json](https://api.clo.ru/openapi.json))
with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen), plus a small
hand-written ergonomic layer (auth, retry, logging, pagination, filtering, error
classification). The generated client always covers 100% of the API and is
regenerated daily, so new endpoints appear automatically with no SDK changes.

> **Upgrading from v2?** v2 hand-wrote a request struct per endpoint
> (`services/...`). v3 replaces all of that with the generated client. See
> [MIGRATION.md](MIGRATION.md). v3 is a separate module path, so nothing breaks
> until you opt in.
>
> **Still on v2?** v2 is frozen but fully available — `/v2` is a distinct module
> path served from its own tags. Pin `github.com/clo-ru/cloapi-go-client/v2@v2.0.0`,
> or track the `v2` branch. The default branch (`main`) now tracks v3.

## Install

```
go get github.com/clo-ru/cloapi-go-client/v3
```

## Usage

Create a client with your JWT token, then call any generated `*WithResponse` method.
On success the typed payload is in `resp.OK`; on a non-2xx response the call returns
an `*ApiError` as the `error` (and also populates `resp.Error`).

```go
package main

import (
	"context"
	"fmt"

	cloapi "github.com/clo-ru/cloapi-go-client/v3"
)

func main() {
	cli, err := cloapi.New("your-jwt-token")
	if err != nil {
		panic(err)
	}

	resp, err := cli.AccountBalanceWithResponse(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp.OK)
}
```

### Options

`New(token, opts...)` accepts functional options (all optional):

| Option | Purpose |
|---|---|
| `WithBaseURL(url)` | Override the API base URL (default `https://api.clo.ru`). |
| `WithTimeout(d)` | HTTP client timeout (default 30s). |
| `WithLogger(*slog.Logger)` | Logger for the logging transport. |
| `WithHTTPClient(*http.Client)` | Supply a custom HTTP client (its transport is wrapped). |
| `WithRetry(count, backoff)` | Retry transient failures (5xx + 429) on idempotent methods. |

```go
cli, _ := cloapi.New(token,
	cloapi.WithTimeout(15*time.Second),
	cloapi.WithRetry(3, 500*time.Millisecond),
)
```

### Errors

A non-2xx response is returned as `*ApiError`. Classify it with the helpers:

```go
resp, err := cli.ServerDetailWithResponse(ctx, serverID)
switch {
case cloapi.IsNotFound(err):
	// 404 — e.g. drop the resource from Terraform state
case cloapi.IsClientError(err):
	// 4xx
case cloapi.IsServerError(err):
	// 5xx
case err != nil:
	return err
}
```

`AsApiError(err)` returns the underlying `*ApiError`; `HasStatus(err, code)` checks a
specific status.

### Query parameters, filtering & ordering

Every method takes variadic `RequestEditorFn` callbacks. Build them with the helpers:

```go
resp, err := cli.ProjectServerListWithResponse(ctx, projectID,
	cloapi.WithPage(50, 0),
	cloapi.WithOrderBy("-created"),
	cloapi.WithFilter(cloapi.FilterField{Field: "ram", Condition: "gte", Value: "1024"}),
)
```

Supported filter conditions: `gt`, `gte`, `lt`, `lte`, `range`, `in`. Use `WithOrderBy("field")`
for ascending and `WithOrderBy("-field")` for descending.

### Pagination

`Paginator[T]` walks a list endpoint. Supply a small closure that calls the endpoint
and returns `(items, count)` from the `{count, result}` envelope:

```go
p := cloapi.NewPaginator(50, func(ctx context.Context, limit, offset int) ([]cloapi.ServerSchema, int, error) {
	resp, err := cli.ProjectServerListWithResponse(
		ctx, projectID, cloapi.WithPage(limit, offset))
	if err != nil {
		return nil, 0, err
	}
	return *resp.OK.Result, resp.OK.Count, nil
})

for p.HasNext() {
	page, err := p.Next(ctx)
	if err != nil {
		return err
	}
	// use page
}
// or: all, err := p.All(ctx)
```

## Development

```
make generate   # fetch spec -> fix.jq -> oapi-codegen -> clo_gen.go, then tidy
make all        # generate + remove the temporary spec
go test ./...   # tests cover the hand-written layer only
```

`clo_gen.go` is generated — **do not edit it by hand**. Ergonomics live in the
hand-written files (`client.go`, `transport.go`, `errors.go`, `filter.go`,
`paginator.go`). Spec defects are patched in `spec/fix.jq`; the goal is to fix them
upstream and shrink `fix.jq` toward empty.