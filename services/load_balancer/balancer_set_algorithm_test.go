package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerSetAlgorithmRequest_Build(t *testing.T) {
	dID := "id"
	b := BalancerSetAlgorithmBody{Algorithm: "ROUND_ROBIN"}
	req := &BalancerSetAlgorithmRequest{BalancerID: dID, Body: b}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(balancerSetAlgorithmEndpoint, mocks.MockUrl, dID), b, t)
}

func TestBalancerSetAlgorithmRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &BalancerSetAlgorithmRequest{BalancerID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerSetAlgorithmRequest{BalancerID: "address_id"},
			},
		},
	)
}

func TestBalancerSetAlgorithmRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerSetAlgorithmRequest{}, t)
}
