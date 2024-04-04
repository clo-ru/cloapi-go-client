package load_balancer

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	balancerChangeMonitorEndpoint = "%s/v2/loadbalancers/%s/healthmonitor"
)

type BalancerChangeMonitorRequest struct {
	clo.Request
	BalancerID string
	Body       BalancerChangeMonitorBody
}

type BalancerChangeMonitorBody struct {
	Delay         int    `json:"delay"`
	Timeout       int    `json:"timeout"`
	MaxRetries    int    `json:"max_retries"`
	Type          string `json:"type"`
	UrlPath       string `json:"url_path"`
	HttpMethod    string `json:"http_method"`
	ExpectedCodes string `json:"expected_codes"`
}

func (r *BalancerChangeMonitorRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *BalancerChangeMonitorRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPut, fmt.Sprintf(balancerChangeMonitorEndpoint, baseUrl, r.BalancerID), authToken, body)
}
