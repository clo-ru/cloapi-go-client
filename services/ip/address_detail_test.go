package ip

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressDetailRequest_BuildRequest(t *testing.T) {
	dID := "address_id"
	req := &AddressDetailRequest{AddressID: dID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(addressDetailEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestAddressDetailRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id":"fipid","status":"ACTIVE","address":"192.168.1.1"}}`, http.StatusOK
				},
				Req: &AddressDetailRequest{AddressID: "id"},
				Expected: &AddressDetailResponse{
					Result: Address{
						ID:      "fipid",
						Status:  "ACTIVE",
						Address: "192.168.1.1",
					},
				},
				Actual: &AddressDetailResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &AddressDetailRequest{AddressID: "id"},
			},
		},
	)
}

func TestAddressDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressDetailRequest{}, t)
}
