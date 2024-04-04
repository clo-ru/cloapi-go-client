package clo

import (
	"context"
	"fmt"
	"strconv"
)

type Paginator struct {
	limit  int
	offset int

	client  *ApiClient
	request RequestInt

	lastPage bool
}

func NewPaginator(client *ApiClient, request RequestInt, limit, offset int) *Paginator {
	return &Paginator{client: client, request: request, limit: limit, offset: offset}
}

func (lp *Paginator) LastPage() bool {
	return lp.lastPage
}

func (lp *Paginator) NextPage(ctx context.Context, dst ListResponseInterface) error {
	if lp.LastPage() {
		return fmt.Errorf("no more pages")
	}
	lp.request.WithQueryParams(QueryParam{"limit": {strconv.Itoa(lp.limit)}, "offset": {strconv.Itoa(lp.offset)}})
	if err := lp.client.DoRequest(ctx, lp.request, dst); err != nil {
		return err
	}
	lp.offset += lp.limit
	if dst.GetCount() <= lp.limit+lp.offset {
		lp.lastPage = true
	}
	return nil
}
