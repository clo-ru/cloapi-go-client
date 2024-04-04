package ip

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestChangerPrimaryRequest_BuildRequest(t *testing.T) {
	dID := "address_id"
	req := &AddressPrimaryChangeRequest{AddressID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(addressPrimaryChangeEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestChangerPrimaryRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &AddressPrimaryChangeRequest{AddressID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &AddressPrimaryChangeRequest{AddressID: "address_id"},
			},
		},
	)
}

func TestChangerPrimaryRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressPrimaryChangeRequest{}, t)
}
