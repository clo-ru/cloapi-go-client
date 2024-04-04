package servers

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo/mocks"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"net/http"
	"testing"
)

func TestServerChangePasswdRequest_BuildRequest(t *testing.T) {
	b := ServerChangePasswdBody{Password: "m"}
	ID := "id"
	req := &ServerChangePasswdRequest{Body: b, ServerID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(serverChangePasswdEndpoint, mocks.MockUrl, ID), b, t)
}

func TestServerChangePasswdRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return "", http.StatusOK
				},
				Req: &ServerChangePasswdRequest{ServerID: "id"},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &ServerChangePasswdRequest{ServerID: "id"},
			},
		},
	)
}

func TestServerChangePasswdRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerChangePasswdRequest{}, t)
}
