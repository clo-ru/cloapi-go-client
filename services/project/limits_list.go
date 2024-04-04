package project

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	limitsListEndpoint = "%s/v2/projects/%s/limits"
)

type LimitsListRequest struct {
	clo.Request
	ProjectID string
}

func (r *LimitsListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*LimitsListResponse, error) {
	resp := &LimitsListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *LimitsListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(limitsListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *LimitsListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *LimitsListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
