package ip

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &AddressListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(addressListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestAddressListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2, "result": [{"id": "first_item_id", "ptr": "host.com", "ddos_protection":true},{"id": "second_item_id", "ptr": "host.com", "ddos_protection":true}]}`, http.StatusOK
				},
				Req: &AddressListRequest{ProjectID: "project_id"},
				Expected: &AddressListResponse{
					Count: 2,
					Result: []Address{
						{ID: "first_item_id", Ptr: "host.com", DdosProtection: true},
						{ID: "second_item_id", Ptr: "host.com", DdosProtection: true},
					},
				},
				Actual: &AddressListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &AddressListRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestAddressListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressListRequest{}, t)
}

func TestAddressListRequest_Filtering(t *testing.T) {
	intTesting.FilterTest(&AddressListRequest{}, t)
}
