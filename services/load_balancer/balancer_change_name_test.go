package load_balancer

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerChangeNameRequest_BuildRequest(t *testing.T) {
	b := BalancerChangeNameBody{Name: "name"}
	dID := "object_id"
	req := &BalancerChangeNameRequest{Body: b, ObjectId: dID}
	intTesting.BuildTest(req, http.MethodPatch, fmt.Sprintf(balancerChangeNameEndpoint, mocks.MockUrl, dID), b, t)
}

func TestBalancerChangeNameRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &BalancerChangeNameRequest{ObjectId: "object_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerChangeNameRequest{ObjectId: "object_id"},
			},
		},
	)
}

func TestBalancerChangeNameRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerChangeNameRequest{}, t)
}
