package sshkeys

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestKeyPairDetailRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &KeyPairDetailRequest{KeypairID: ID}

	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(keypairDetailEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestKeyPairDetailRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"name":"kp_name","id":"id","public_key":"pubkey"}}`, http.StatusOK
				},
				Req: &KeyPairDetailRequest{KeypairID: "id"},
				Expected: &KeyPairDetailResponse{
					Result: KeyPair{
						ID:        "id",
						Name:      "kp_name",
						PublicKey: "pubkey",
					},
				},
				Actual: &KeyPairDetailResponse{},
			},
		},
	)
}

func TestKeyPairDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&KeyPairDetailRequest{}, t)
}
