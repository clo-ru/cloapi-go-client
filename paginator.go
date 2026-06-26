package cloapi

import (
	"context"
	"fmt"
)

// PageFunc fetches one page given a limit and offset, returning the page items and
// the total count reported by the API's {count, result} envelope.
type PageFunc[T any] func(ctx context.Context, limit, offset int) (items []T, count int, err error)

// Paginator walks a list endpoint page by page. It is generic over the element type
// and driven by a PageFunc, so it works for every list endpoint (present and future)
// without per-endpoint code. The PageFunc is a thin wrapper around a generated
// *WithResponse call, e.g.:
//
//	p := cloapi.NewPaginator(50, func(ctx context.Context, limit, offset int) ([]cloapi.Server, int, error) {
//		resp, err := cli.ProjectServersListV2ProjectsObjectIdServersGetWithResponse(
//			ctx, projectID, cloapi.WithPage(limit, offset))
//		if err != nil {
//			return nil, 0, err
//		}
//		return *resp.OK.Result, resp.OK.Count, nil
//	})
//	for p.HasNext() {
//		page, err := p.Next(ctx)
//		...
//	}
type Paginator[T any] struct {
	limit    int
	offset   int
	count    int
	started  bool
	lastPage bool
	fetch    PageFunc[T]
}

// NewPaginator creates a Paginator with the given page size and fetch function.
func NewPaginator[T any](limit int, fetch PageFunc[T]) *Paginator[T] {
	if limit <= 0 {
		limit = 50
	}
	return &Paginator[T]{limit: limit, fetch: fetch}
}

// HasNext reports whether another page is available. It is true until a page has
// been fetched that reaches or exceeds the reported total count.
func (p *Paginator[T]) HasNext() bool {
	return !p.lastPage
}

// Next fetches the next page and advances the offset. It returns an error if called
// after the last page has been consumed.
func (p *Paginator[T]) Next(ctx context.Context) ([]T, error) {
	if p.lastPage {
		return nil, fmt.Errorf("cloapi: no more pages")
	}
	items, count, err := p.fetch(ctx, p.limit, p.offset)
	if err != nil {
		return nil, err
	}
	p.started = true
	p.count = count
	p.offset += p.limit
	if p.offset >= count {
		p.lastPage = true
	}
	return items, nil
}

// Count returns the total item count reported by the API. It is only meaningful
// after the first Next call.
func (p *Paginator[T]) Count() int {
	return p.count
}

// All drains every remaining page and returns the concatenated items.
func (p *Paginator[T]) All(ctx context.Context) ([]T, error) {
	var all []T
	for p.HasNext() {
		page, err := p.Next(ctx)
		if err != nil {
			return all, err
		}
		all = append(all, page...)
	}
	return all, nil
}
