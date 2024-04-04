package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	balancerStartEndpoint = "%s/v2/loadbalancers/%s/start"
)

type BalancerStartRequest struct {
	clo.Request
	BalancerID string
}

func (r *BalancerStartRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *BalancerStartRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(balancerStartEndpoint, baseUrl, r.BalancerID), authToken, nil)
}
