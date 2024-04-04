package sshkeys

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestKeyPairDeleteRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &KeyPairDeleteRequest{KeypairID: ID}
	intTesting.BuildTest(req, http.MethodDelete, fmt.Sprintf(keypairDeleteEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestKeyPairDeleteRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return "1",
						http.StatusOK
				},
				Req: &KeyPairDeleteRequest{KeypairID: "id"},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &KeyPairDeleteRequest{KeypairID: "id"},
			},
		},
	)
}

func TestKeyPairDeleteRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&KeyPairDeleteRequest{}, t)
}
