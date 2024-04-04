package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerDetailRequest_Build(t *testing.T) {
	dID := "id"
	req := &BalancerDetailRequest{ObjectId: dID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(balancerDetailEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestBalancerDetailRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id":"id1","algorithm":"algo","addresses":["ids"],"healthmonitor":{"http_method":"get"}}}`, http.StatusOK
				},
				Req: &BalancerDetailRequest{ObjectId: "id"},
				Expected: &LoadBalancerDetailResponse{
					Result: LoadBalancer{
						ID:            "id1",
						Algorithm:     "algo",
						Addresses:     []string{"ids"},
						HealthMonitor: BalancerMonitor{HttpMethod: "get"},
					},
				},
				Actual: &LoadBalancerDetailResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &BalancerDetailRequest{ObjectId: "id"},
			},
		},
	)
}

func TestBalancerDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerDetailRequest{}, t)
}
