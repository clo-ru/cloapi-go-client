package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerStartRequest_Build(t *testing.T) {
	dID := "id"
	req := &BalancerStartRequest{BalancerID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(balancerStartEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestBalancerStartRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &BalancerStartRequest{BalancerID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerStartRequest{BalancerID: "address_id"},
			},
		},
	)
}

func TestBalancerStartRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerStartRequest{}, t)
}
