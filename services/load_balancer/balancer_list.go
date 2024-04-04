package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	balancerListEndpoint = "%s/v2/projects/%s/loadbalancers"
)

type BalancerListRequest struct {
	clo.Request
	ProjectID string
}

func (r *BalancerListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*LoadBalancerListResponse, error) {
	resp := &LoadBalancerListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}
func (r *BalancerListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(balancerListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}
func (r *BalancerListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *BalancerListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
