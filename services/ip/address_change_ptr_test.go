package ip

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressPtrChangeRequest_BuildRequest_BuildRequest(t *testing.T) {
	b := AddressPtrChangeBody{Value: "val"}
	dID := "address_id"
	req := &AddressPtrChangeRequest{Body: b, AddressID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(addressPtrChangeEndpoint, mocks.MockUrl, dID), b, t)
}

func TestAddressPtrChangeRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &AddressPtrChangeRequest{AddressID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &AddressPtrChangeRequest{AddressID: "address_id"},
			},
		},
	)
}

func TestAddressPtrChangeRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressPtrChangeRequest{}, t)
}
