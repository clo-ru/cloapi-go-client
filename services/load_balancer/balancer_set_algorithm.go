package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	balancerSetAlgorithmEndpoint = "%s/v2/loadbalancers/%s/algorithm"
)

type BalancerSetAlgorithmRequest struct {
	clo.Request
	BalancerID string
	Body       BalancerSetAlgorithmBody
}

type BalancerSetAlgorithmBody struct {
	Algorithm          string `json:"algorithm"`
	SessionPersistence bool   `json:"session_persistence"`
}

func (r *BalancerSetAlgorithmRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *BalancerSetAlgorithmRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(balancerSetAlgorithmEndpoint, baseUrl, r.BalancerID), authToken, body)
}
