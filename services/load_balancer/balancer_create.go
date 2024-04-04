package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	balancerCreateEndpoint = "%s/v2/projects/%s/loadbalancers"
)

type BalancerCreateRequest struct {
	clo.Request
	ProjectID string
	Body      BalancerCreateBody
}

type BalancerCreateBody struct {
	Name               string              `json:"name"`
	Algorithm          string              `json:"algorithm"`
	SessionPersistence bool                `json:"session_persistence"`
	Address            BalancerBodyAddress `json:"address,omitempty"`
	HealthMonitor      BalancerBodyMonitor `json:"healthmonitor"`
	Rules              []BalancerRuleBody  `json:"rules"`
}

type BalancerBodyAddress struct {
	ID             string `json:"id"`
	DdosProtection bool   `json:"ddos_protection"`
}

type BalancerBodyMonitor struct {
	Delay         int    `json:"delay"`
	Timeout       int    `json:"timeout"`
	MaxRetries    int    `json:"max_retries"`
	Type          string `json:"type"`
	UrlPath       string `json:"url_path,omitempty"`
	HttpMethod    string `json:"http_method,omitempty"`
	ExpectedCodes string `json:"expected_codes,omitempty"`
}

type BalancerRuleBody struct {
	AddressId            string `json:"address_id"`
	ExternalProtocolPort int    `json:"external_protocol_port"`
	InternalProtocolPort int    `json:"internal_protocol_port"`
}

func (r *BalancerCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *BalancerCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(balancerCreateEndpoint, baseUrl, r.ProjectID), authToken, body)
}
