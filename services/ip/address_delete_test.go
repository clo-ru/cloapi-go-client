package ip

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressDeleteRequest_BuildRequest(t *testing.T) {
	dID := "address_id"
	req := &AddressDeleteRequest{AddressID: dID}
	intTesting.BuildTest(req, http.MethodDelete, fmt.Sprintf(addressDeleteEndpoint, mocks.MockUrl, dID), nil, t)
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
				Req: &AddressDeleteRequest{AddressID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &AddressDeleteRequest{AddressID: "address_id"},
			},
		},
	)
}

func TestAddressDeleteRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressDeleteRequest{}, t)
}
