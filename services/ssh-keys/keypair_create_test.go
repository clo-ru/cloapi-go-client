package sshkeys

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestKeyPairCreateRequest_BuildRequest(t *testing.T) {
	projectId := "id"
	b := KeyPairCreateBody{}
	req := &KeyPairCreateRequest{Body: b, ProjectID: projectId}

	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(keypairListEndpoint, mocks.MockUrl, projectId), b, t)
}

func TestKeyPairCreateRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return `{"result":{"id":"key_id"}}`, http.StatusOK },
				Req:            &KeyPairCreateRequest{},
				Expected:       &clo.ResponseCreated{Result: clo.IdResult{ID: "key_id"}},
				Actual:         &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &KeyPairCreateRequest{},
				Expected:       nil,
			},
		},
	)
}

func TestKeyPairCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&KeyPairCreateRequest{}, t)
}
