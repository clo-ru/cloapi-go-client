package load_balancer

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressCreateRequest_BuildRequest(t *testing.T) {
	b := BalancerRulesCreateBody{AddressID: "addr"}
	dID := "id"
	req := &BalancerRulesCreateRequest{Body: b, BalancerID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(balancerRulesCreateEndpoint, mocks.MockUrl, dID), b, t)
}

func TestAddressCreateRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id":"disk_id"}}`, http.StatusOK
				},
				Req:      &BalancerRulesCreateRequest{BalancerID: "project_id"},
				Expected: &clo.ResponseCreated{Result: clo.IdResult{ID: "disk_id"}},
				Actual:   &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerRulesCreateRequest{BalancerID: "project_id"},
			},
		},
	)
}

func TestAddressCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerRulesCreateRequest{}, t)
}
