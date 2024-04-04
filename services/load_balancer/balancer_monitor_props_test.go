package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerChangeMonitorRequest_Build(t *testing.T) {
	b := BalancerChangeMonitorBody{Delay: 10}
	dID := "object_id"
	req := &BalancerChangeMonitorRequest{Body: b, BalancerID: dID}
	intTesting.BuildTest(req, http.MethodPut, fmt.Sprintf(balancerChangeMonitorEndpoint, mocks.MockUrl, dID), b, t)
}

func TestBalancerChangeMonitorRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &BalancerChangeMonitorRequest{BalancerID: "object_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerChangeMonitorRequest{BalancerID: "object_id"},
			},
		},
	)
}

func TestBalancerChangeMonitorRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerChangeMonitorRequest{}, t)
}
