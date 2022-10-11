package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	serverListEndpoint = "/v1/projects/%s/servers"
)

type ServerListRequest struct {
	clo.Request
	ProjectID string
}

func (r *ServerListRequest) Make(ctx context.Context, cli *clo.ApiClient) (ServerListResponse, error) {
	rawReq, e := r.buildRequest(ctx, cli.Options)
	if e != nil {
		return ServerListResponse{}, e
	}
	rawResp, requestError := r.MakeRequest(rawReq, cli)
	if requestError != nil {
		return ServerListResponse{}, requestError
	}
	defer rawResp.Body.Close()
	var resp ServerListResponse
	if e = resp.FromJsonBody(rawResp.Body); e != nil {
		return ServerListResponse{}, e
	}
	return resp, nil
}

func (r *ServerListRequest) buildRequest(ctx context.Context, cliOptions map[string]interface{}) (*http.Request, error) {
	authKey, ok := cliOptions["auth_key"].(string)
	if !ok {
		return nil, fmt.Errorf("auth_key client options should be a string, %T got", authKey)
	}
	baseUrl, ok := cliOptions["base_url"].(string)
	if !ok {
		return nil, fmt.Errorf("base_url client options should be a string, %T got", baseUrl)
	}
	baseUrl += fmt.Sprintf(serverListEndpoint, r.ProjectID)
	rawReq, e := http.NewRequestWithContext(
		ctx, http.MethodGet, baseUrl, nil,
	)
	if e != nil {
		return nil, e
	}
	h := http.Header{}
	h.Add("Authorization", fmt.Sprintf("Bearer %s", authKey))
	r.WithHeaders(h)
	return rawReq, nil
}

func (r *ServerListRequest) OrderBy(of string) *ServerListRequest {
	r.WithQueryParams(map[string][]string{"order": {of}})
	return r
}

type FilteringField struct {
	FieldName string
	Condition string
	Value     string
}

func (r *ServerListRequest) FilterBy(ff FilteringField) *ServerListRequest {
	switch ff.Condition {
	case "gt", "gte", "lt", "lte", "range", "in":
		condString := fmt.Sprintf("%s__%s", ff.FieldName, ff.Condition)
		r.WithQueryParams(map[string][]string{condString: {ff.Value}})
	}
	return r
}

type PaginatorOptions struct {
	Limit  int
	Offset int
}

type ListPaginator struct {
	op       PaginatorOptions
	client   *clo.ApiClient
	params   ServerListRequest
	lastPage bool
}

func (lp *ListPaginator) LastPage() bool {
	return lp.lastPage
}

func NewListPaginator(client *clo.ApiClient, params ServerListRequest, op PaginatorOptions) (*ListPaginator, error) {
	if op.Limit == 0 {
		return nil, fmt.Errorf("op.Limit should not be 0")
	}
	lp := ListPaginator{
		client: client,
		params: params,
		op:     op,
	}
	return &lp, nil
}

func (lp *ListPaginator) NextPage(ctx context.Context) (ServerListResponse, error) {
	if lp.LastPage() {
		return ServerListResponse{}, fmt.Errorf("no more pages")
	}
	lp.params.WithQueryParams(map[string][]string{"limit": {fmt.Sprintf("%d", lp.op.Limit)}})
	lp.params.WithQueryParams(map[string][]string{"offset": {fmt.Sprintf("%d", lp.op.Offset)}})
	r, e := lp.params.Make(ctx, lp.client)
	if e != nil {
		return ServerListResponse{}, e
	}
	lp.op.Offset += lp.op.Limit
	if r.Count <= lp.op.Limit+lp.op.Offset {
		lp.lastPage = true
	}
	return r, e
}
