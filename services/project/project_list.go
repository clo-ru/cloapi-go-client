package project

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	projectListEndpoint = "%s/v2/projects"
)

type ProjectListRequest struct {
	clo.Request
}

func (r *ProjectListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*ProjectListResponse, error) {
	resp := &ProjectListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *ProjectListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(projectListEndpoint, baseUrl), authToken, nil)
}

func (r *ProjectListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *ProjectListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
