package ip

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressAttachRequest_BuildRequest(t *testing.T) {
	b := AddressAttachBody{ID: "sid", Entity: "server"}
	dID := "address_id"
	req := &AddressAttachRequest{Body: b, AddressID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(addressAttachEndpoint, mocks.MockUrl, dID), b, t)
}

func TestAddressAttachRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &AddressAttachRequest{AddressID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &AddressAttachRequest{AddressID: "address_id"},
			},
		},
	)
}

func TestAddressAttachRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressAttachRequest{}, t)
}
