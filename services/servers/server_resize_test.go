package servers

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestServerResizeRequest_BuildRequest(t *testing.T) {
	ID := "id"
	body := ServerResizeBody{Ram: 1, Vcpus: 2}
	req := &ServerResizeRequest{ServerID: ID, Body: body}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(serverResizeEndpoint, mocks.MockUrl, ID), body, t)

}

func TestServerResizeRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerResizeRequest{}, t)
}

func TestServerResizeRequest_Make(t *testing.T) {
	cases := []intTesting.DoTestCase{
		{
			Name:           "Success",
			BodyStringFunc: func() (string, int) { return "1", http.StatusAccepted },
			Req:            &ServerResizeRequest{ServerID: "id"},
		},
		{
			Name:       "Error",
			ShouldFail: true,
			CheckError: true,
			BodyStringFunc: func() (string, int) {
				return "", http.StatusInternalServerError
			},
			Req: &ServerResizeRequest{ServerID: "id"},
		},
	}
	intTesting.TestDoRequestCases(t, cases)
}
