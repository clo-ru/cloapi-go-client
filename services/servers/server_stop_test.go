package servers

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestServerStopRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &ServerStopRequest{ServerID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(serverStopEndpoint, mocks.MockUrl, ID), nil, t)

}

func TestServerStopRequest_Make(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerStopRequest{}, t)
}

func TestServerStopRequest_MakeRetry(t *testing.T) {
	cases := []intTesting.DoTestCase{
		{
			Name: "Success",
			BodyStringFunc: func() (string, int) {
				return "1", http.StatusAccepted
			},
			Req: &ServerStopRequest{ServerID: "id"},
		},
		{
			Name:       "Error",
			ShouldFail: true,
			CheckError: true,
			BodyStringFunc: func() (string, int) {
				return "", http.StatusInternalServerError
			},
			Req: &ServerStopRequest{ServerID: "id"},
		},
	}
	intTesting.TestDoRequestCases(t, cases)
}
