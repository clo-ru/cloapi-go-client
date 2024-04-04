package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	balancerDeleteEndpoint = "%s/v2/loadbalancers/%s"
)

type BalancerDeleteRequest struct {
	clo.Request
	ObjectId string
}

func (r *BalancerDeleteRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *BalancerDeleteRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodDelete, fmt.Sprintf(balancerDeleteEndpoint, baseUrl, r.ObjectId), authToken, nil)
}
