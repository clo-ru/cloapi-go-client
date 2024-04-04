package servers

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestServerStartRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &ServerStartRequest{ServerID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(serverStartEndpoint, mocks.MockUrl, ID), nil, t)

}

func TestServerStartRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerStartRequest{}, t)
}

func TestServerStartRequest_Make(t *testing.T) {
	cases := []intTesting.DoTestCase{
		{
			Name: "Success",
			BodyStringFunc: func() (string, int) {
				return "1", http.StatusAccepted
			},
			Req: &ServerStartRequest{ServerID: "id"},
		},
		{
			Name:       "Error",
			ShouldFail: true,
			CheckError: true,
			BodyStringFunc: func() (string, int) {
				return "", http.StatusInternalServerError
			},
			Req: &ServerStartRequest{ServerID: "id"},
		},
	}
	intTesting.TestDoRequestCases(t, cases)
}
