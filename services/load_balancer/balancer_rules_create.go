package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	balancerRulesCreateEndpoint = "%s/v2/loadbalancers/%s/rules"
)

type BalancerRulesCreateRequest struct {
	clo.Request
	BalancerID string
	Body       BalancerRulesCreateBody
}

type BalancerRulesCreateBody struct {
	AddressID            string `json:"address_id"`
	ExternalProtocolPort int    `json:"external_protocol_port"`
	InternalProtocolPort int    `json:"internal_protocol_port"`
}

func (r *BalancerRulesCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *BalancerRulesCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(balancerRulesCreateEndpoint, baseUrl, r.BalancerID), authToken, body)
}
