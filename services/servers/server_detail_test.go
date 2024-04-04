package servers

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo/mocks"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"net/http"
	"testing"
)

func TestServerDetailRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &ServerDetailRequest{ServerID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(serverDetailEndpoint, mocks.MockUrl, ID), nil, t)

}

func TestServerDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerDetailRequest{}, t)
}

func TestServerDetailRequest_Make(t *testing.T) {
	var cases = []intTesting.DoTestCase{
		{
			Name:           "Success",
			BodyStringFunc: func() (string, int) { return `{"result":{"id":"sid","flavor":{"ram":2,"vcpus":3}}}`, http.StatusOK },
			Req:            &ServerDetailRequest{ServerID: "id"},
			Expected:       &ServerDetailResponse{Result: Server{ID: "sid", Flavor: ServerFlavor{Ram: 2, Vcpus: 3}}},
			Actual:         &ServerDetailResponse{},
		},
		{
			Name:           "Error",
			ShouldFail:     true,
			CheckError:     true,
			BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
			Req:            &ServerDetailRequest{ServerID: "id"},
		},
	}
	intTesting.TestDoRequestCases(t, cases)
}
