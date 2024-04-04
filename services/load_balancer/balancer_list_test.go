package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &BalancerListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(balancerListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestBalancerListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count":1,"result":[{"id":"id1","algorithm":"algo","addresses":["ids"],"healthmonitor":{"http_method":"get"}}]}`, http.StatusOK
				},
				Req: &BalancerListRequest{ProjectID: "project_id"},
				Expected: &LoadBalancerListResponse{
					Count: 1,
					Result: []LoadBalancer{
						{
							ID:            "id1",
							Algorithm:     "algo",
							Addresses:     []string{"ids"},
							HealthMonitor: BalancerMonitor{HttpMethod: "get"},
						},
					},
				},
				Actual: &LoadBalancerListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerListRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestBalancerListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerListRequest{}, t)
}

func TestBalancerListRequest_Filtering(t *testing.T) {
	intTesting.FilterTest(&BalancerListRequest{}, t)
}
