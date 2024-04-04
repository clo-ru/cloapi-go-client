package load_balancer

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestBalancerCreateRequest_Build(t *testing.T) {
	b := BalancerCreateBody{Name: "test"}
	projectID := "project_id"
	req := &BalancerCreateRequest{Body: b, ProjectID: projectID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(balancerCreateEndpoint, mocks.MockUrl, projectID), b, t)
}

func TestBalancerCreateRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id":"disk_id"}}`, http.StatusOK
				},
				Req:      &BalancerCreateRequest{ProjectID: "project_id"},
				Expected: &clo.ResponseCreated{Result: clo.IdResult{ID: "disk_id"}},
				Actual:   &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &BalancerCreateRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestBalancerCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&BalancerCreateRequest{}, t)
}
