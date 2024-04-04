package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	balancerChangeNameEndpoint = "%s/v2/loadbalancers/%s"
)

type BalancerChangeNameRequest struct {
	clo.Request
	ObjectId string
	Body     BalancerChangeNameBody
}

type BalancerChangeNameBody struct {
	Name string `json:"name"`
}

func (r *BalancerChangeNameRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *BalancerChangeNameRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPatch, fmt.Sprintf(balancerChangeNameEndpoint, baseUrl, r.ObjectId), authToken, body)
}
