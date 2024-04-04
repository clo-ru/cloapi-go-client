package storage

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3KeysResetRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &S3KeysResetRequest{UserID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(s3KeysResetEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestS3KeysResetRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{{
			Name:           "Success",
			BodyStringFunc: func() (string, int) { return `{"result":{"access_key":"aKey","secret_key":"secKey"}}`, http.StatusOK },
			Req:            &S3KeysResetRequest{},
			Expected:       &S3KeysResetResponse{Result: S3Key{AccessKey: "aKey", SecretKey: "secKey"}},
			Actual:         &S3KeysResetResponse{},
		},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3KeysResetRequest{},
			}},
	)

}

func TestS3KeysResetRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3KeysResetRequest{}, t)
}
