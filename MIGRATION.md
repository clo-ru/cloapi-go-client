# Migrating from v2 to v3

v3 is a ground-up change: the per-endpoint request structs under `services/*` are
gone, replaced by a client **generated from the CLO OpenAPI spec**. You now call
generated `*WithResponse` methods on a single client instead of building a request
struct per call.

The `/v2` import path is unaffected and keeps working at its existing tags — v3 lives
at a new module path (`.../v3`), so you migrate on your own schedule. **v2 is frozen:
it receives no further changes.**

## Import path

```go
// v2
import (
	"github.com/clo-ru/cloapi-go-client/v2/clo"
	"github.com/clo-ru/cloapi-go-client/v2/services/servers"
)

// v3
import cloapi "github.com/clo-ru/cloapi-go-client/v3"
```

Everything is now in one package; `clo/request_tools` and `services/*` no longer exist.

## Mapping

| v2 | v3 |
|---|---|
| `clo.NewDefaultClient(token, url)` → `*clo.ApiClient` | `cloapi.New(token, cloapi.WithBaseURL(url))` → `*cloapi.ClientWithResponses` |
| `services/x.XRequest{...}.Do(ctx, cli)` | `cli.<Op>WithResponse(ctx, ...pathParams, body, editors...)` |
| `resp.Result.<field>` | `resp.OK.<field>` |
| `services/x.XBody{...}` | generated `<Op>JSONRequestBody{...}` |
| `req.WithRetry(n, d)` (per request) | `cloapi.WithRetry(n, d)` (once, on the client) |
| `req.WithLog(logger)` | `cloapi.WithLogger(*slog.Logger)` (once, on the client) |
| `req.WithQueryParams(m)` | `cloapi.WithQuery(m)` editor |
| `clo.AddFilterToRequest`, `FilteringField` | `cloapi.WithFilter(cloapi.FilterField{...})` editor |
| `req.OrderBy("field")` | `cloapi.WithOrderBy("field")` editor |
| `errors.As(err, &clo_tools.DefaultError{}) && .Code == 404` | `cloapi.IsNotFound(err)` |
| `clo.NewPaginator(cli, req, limit, offset)` + `NextPage`/`LastPage` | `cloapi.NewPaginator(limit, fetchFn)` + `HasNext`/`Next`/`All` |

## Error handling

In v2 a non-2xx returned a `clo_tools.DefaultError`. In v3 it returns an `*ApiError`
(and also populates `resp.Error`). Classify with the helpers:

```go
// v2
resp, err := req.Do(ctx, cli)
apiErr := clo_tools.DefaultError{}
if errors.As(err, &apiErr) && apiErr.Code == 404 {
	// gone
}

// v3
resp, err := cli.ServerDetailWithResponse(ctx, id)
if cloapi.IsNotFound(err) {
	// gone
}
```

## Example: create a server

```go
// v2
req := servers.ServerCreateRequest{
	ProjectID: projectID,
	Body: servers.ServerCreateBody{
		Name:   "my_server",
		Image:  "5f3e...image-uuid", // image ID
		Flavor: servers.ServerFlavorBody{Ram: 2048, Vcpus: 2},
	},
}
req.WithRetry(2, 0)
resp, err := req.Do(ctx, cli)
id := resp.Result.ID

// v3
cli, _ := cloapi.New(token, cloapi.WithRetry(2, 0))

imageID := "5f3e...image-uuid"
body := cloapi.ServerCreateJSONRequestBody{
	Name:  "my_server",
	Image: &imageID, // image ID (not a name/slug); optional fields are pointers
}
body.Flavor.Ram = 2048 // Flavor is an inlined struct
body.Flavor.Vcpus = 2

resp, err := cli.ServerCreateWithResponse(ctx, projectID, body)
if err != nil {
	return err
}
id := resp.OK.Result.Id
```

## A note on generated shapes

Method and type names are clean (`ServerCreateWithResponse`,
`ServerCreateJSONRequestBody`), derived from stable `operationId`s in the spec. Two
shape artifacts remain that come from the spec, not the SDK: optional fields are
pointers, and a few request bodies use inlined anonymous structs (e.g. server-create
`Flavor`). These improve as the spec's schemas are cleaned up upstream — with no SDK
code change, just a regeneration. Pin a tag if you need name/shape stability.