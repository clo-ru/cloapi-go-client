package servers

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestServerCreateRequest_BuildRequest(t *testing.T) {
	b := ServerCreateBody{Name: "m"}
	ID := "id"
	req := &ServerCreateRequest{Body: b, ProjectID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(serverCreateEndpoint, mocks.MockUrl, ID), b, t)
}

func TestServerCreateRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return `{"result":{"id":"sid"}}`, http.StatusOK },
				Req:            &ServerCreateRequest{ProjectID: "id"},
				Expected:       &clo.ResponseCreated{Result: clo.IdResult{ID: "sid"}},
				Actual:         &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &ServerCreateRequest{ProjectID: "id"},
			},
		},
	)

}

func TestServerCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ServerCreateRequest{ProjectID: "id"}, t)
}
