package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerDeleteRequest_BuildRequest(t *testing.T) {
	dID := "object_id"
	req := &BalancerDeleteRequest{ObjectId: dID}
	intTesting.BuildTest(req, http.MethodDelete, fmt.Sprintf(balancerDeleteEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestAddressDeleteRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return "", http.StatusOK
				},
				Req: &BalancerDeleteRequest{ObjectId: "object_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerDeleteRequest{ObjectId: "object_id"},
			},
		},
	)
}

func TestAddressDeleteRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerDeleteRequest{}, t)
}
