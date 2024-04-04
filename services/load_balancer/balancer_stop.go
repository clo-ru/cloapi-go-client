package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	balancerStopEndpoint = "%s/v2/loadbalancers/%s/stop"
)

type BalancerStopRequest struct {
	clo.Request
	BalancerID string
}

func (r *BalancerStopRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *BalancerStopRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(balancerStopEndpoint, baseUrl, r.BalancerID), authToken, nil)
}
