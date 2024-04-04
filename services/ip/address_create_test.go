package ip

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressCreateRequest_BuildRequest(t *testing.T) {
	b := AddressCreateBody{DdosProtection: true}
	dID := "volume_id"
	req := &AddressCreateRequest{Body: b, ProjectID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(addressCreateEndpoint, mocks.MockUrl, dID), b, t)
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
				Req:      &AddressCreateRequest{ProjectID: "project_id"},
				Expected: &clo.ResponseCreated{Result: clo.IdResult{ID: "disk_id"}},
				Actual:   &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &AddressCreateRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestAddressCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressCreateRequest{}, t)
}
