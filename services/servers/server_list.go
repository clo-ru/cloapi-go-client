package servers

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	serverListEndpoint = "%s/v2/projects/%s/servers"
)

type ServerListRequest struct {
	clo.Request
	ProjectID string
}

func (r *ServerListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*ServerListResponse, error) {
	resp := &ServerListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *ServerListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(serverListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *ServerListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *ServerListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
