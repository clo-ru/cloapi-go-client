package load_balancer

import "github.com/clo-ru/cloapi-go-client/clo"

type LoadBalancer struct {
	ID                 string          `json:"id"`
	Project            string          `json:"project"`
	Name               string          `json:"name"`
	Algorithm          string          `json:"algorithm"`
	Status             string          `json:"status"`
	CreatedIn          string          `json:"created_in"`
	UpdatedIn          string          `json:"updated_in"`
	RulesCount         int             `json:"rules_count"`
	SessionPersistence bool            `json:"session_persistence"`
	HealthMonitor      BalancerMonitor `json:"healthmonitor"`
	Addresses          []string        `json:"addresses"`
}

type BalancerMonitor struct {
	Type          string `json:"type"`
	UrlPath       string `json:"url_path"`
	HttpMethod    string `json:"http_method"`
	ExpectedCodes string `json:"expected_codes"`
	Delay         int    `json:"delay"`
	Timeout       int    `json:"timeout"`
	MaxRetries    int    `json:"max_retries"`
}

type BalancerRule struct {
	ID                   string `json:"id"`
	ExternalProtocolPort int    `json:"external_protocol_port"`
	InternalProtocolPort int    `json:"internal_protocol_port"`
	Server               string `json:"server"`
	Loadbalancer         string `json:"loadbalancer"`
	Status               string `json:"status"`
	Address              string `json:"address"`
}

type LoadBalancerDetailResponse = clo.Response[LoadBalancer]
type LoadBalancerListResponse = clo.ListResponse[LoadBalancer]

type BalancerRuleListResponse = clo.ListResponse[BalancerRule]
