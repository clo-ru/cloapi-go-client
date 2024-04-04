package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerStopRequest_Build(t *testing.T) {
	dID := "id"
	req := &BalancerStopRequest{BalancerID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(balancerStopEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestBalancerStopRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &BalancerStopRequest{BalancerID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerStopRequest{BalancerID: "address_id"},
			},
		},
	)
}

func TestBalancerStopRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerStopRequest{}, t)
}
