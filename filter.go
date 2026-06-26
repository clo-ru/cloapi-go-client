package cloapi

import (
	"context"
	"net/http"
	"strconv"
)

// The generated client takes no typed query-parameter objects (the spec documents
// none), but every *WithResponse method accepts variadic RequestEditorFn callbacks.
// These helpers build those callbacks, so query/filter/order/pagination params apply
// uniformly to any endpoint — present or future. Example:
//
//	resp, err := cli.ProjectListV2ProjectsGetWithResponse(ctx,
//		cloapi.WithPage(50, 0),
//		cloapi.WithFilter(cloapi.FilterField{Field: "created", Condition: "gte", Value: "2024-01-01"}),
//	)

// WithQuery adds raw query parameters to the request.
func WithQuery(params map[string]string) RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		return nil
	}
}

// WithPage adds limit/offset pagination query parameters.
func WithPage(limit, offset int) RequestEditorFn {
	return WithQuery(map[string]string{
		"limit":  strconv.Itoa(limit),
		"offset": strconv.Itoa(offset),
	})
}

// WithOrderBy adds an ordering query parameter.
func WithOrderBy(field string) RequestEditorFn {
	return WithQuery(map[string]string{"ordering": field})
}

// FilterField describes a single field filter using the API's `field__condition`
// query convention.
type FilterField struct {
	Field     string
	Condition string
	Value     string
}

// supportedConditions mirrors the conditions the API accepts on filter fields.
var supportedConditions = map[string]bool{
	"gt": true, "gte": true, "lt": true, "lte": true, "range": true, "in": true,
}

// WithFilter adds one or more field filters. Unsupported conditions are skipped.
func WithFilter(filters ...FilterField) RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		q := req.URL.Query()
		for _, f := range filters {
			if !supportedConditions[f.Condition] {
				continue
			}
			q.Add(f.Field+"__"+f.Condition, f.Value)
		}
		req.URL.RawQuery = q.Encode()
		return nil
	}
}
