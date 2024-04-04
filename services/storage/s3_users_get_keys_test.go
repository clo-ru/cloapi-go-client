package storage

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3KeysGetRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &S3KeysGetRequest{UserID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(s3KeysGetEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestS3KeysGetRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":[{"access_key":"aKey","secret_key":"secKey"}], "count": 1}`, http.StatusOK
				},
				Req:      &S3KeysGetRequest{},
				Expected: &S3KeysGetResponse{Result: []S3Key{{AccessKey: "aKey", SecretKey: "secKey"}}, Count: 1},
				Actual:   &S3KeysGetResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3KeysGetRequest{},
			}},
	)
}

func TestS3KeysGetRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3KeysGetRequest{}, t)
}
