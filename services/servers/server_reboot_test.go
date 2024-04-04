package servers

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestServerRebootRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &ServerRebootRequest{ServerID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(serverRebootEndpoint, mocks.MockUrl, ID), nil, t)

}

func TestServerRebootRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerRebootRequest{}, t)
}

func TestServerRebootRequest_Make(t *testing.T) {
	cases := []intTesting.DoTestCase{
		{
			Name:           "Success",
			BodyStringFunc: func() (string, int) { return "1", http.StatusOK },
			Req:            &ServerRebootRequest{ServerID: "id"},
		},
		{
			Name:           "Error",
			ShouldFail:     true,
			CheckError:     true,
			BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
			Req:            &ServerRebootRequest{ServerID: "id"},
		},
	}
	intTesting.TestDoRequestCases(t, cases)
}
