package servers

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestServerDeleteRequest_BuildRequest(t *testing.T) {
	ID := "id"
	b := ServerDeleteBody{ClearFstab: true}
	req := &ServerDeleteRequest{ServerID: ID, Body: b}
	intTesting.BuildTest(req, http.MethodDelete, fmt.Sprintf(serverDeleteEndpoint, mocks.MockUrl, ID), b, t)
}

func TestServerDeleteRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "1", http.StatusOK },
				Req:            &ServerDeleteRequest{ServerID: "id"},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &ServerDeleteRequest{ServerID: "id"},
			},
		},
	)
}

func TestServerDeleteRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerDeleteRequest{ServerID: "id"}, t)
}
