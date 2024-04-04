package ip

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestAddressDetachRequest_Build(t *testing.T) {
	dID := "address_id"
	req := &AddressDetachRequest{AddressID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(addressDetachEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestAddressDetachRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &AddressDetachRequest{AddressID: "address_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &AddressDetachRequest{AddressID: "address_id"},
			},
		},
	)
}

func TestAddressDetachRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&AddressDetachRequest{}, t)
}
