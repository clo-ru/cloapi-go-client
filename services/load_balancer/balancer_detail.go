package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	balancerDetailEndpoint = "%s/v2/loadbalancers/%s/detail"
)

type BalancerDetailRequest struct {
	clo.Request
	ObjectId string
}

func (r *BalancerDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*LoadBalancerDetailResponse, error) {
	res := &LoadBalancerDetailResponse{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *BalancerDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(balancerDetailEndpoint, baseUrl, r.ObjectId), authToken, nil)
}
