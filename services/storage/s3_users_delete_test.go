package storage

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3UserDeleteRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &S3UserDeleteRequest{
		UserID: ID,
	}
	intTesting.BuildTest(req, http.MethodDelete, fmt.Sprintf(s3UserDeleteEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestS3UserDeleteRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{{
			Name:           "Success",
			BodyStringFunc: func() (string, int) { return "1", http.StatusOK },
			Req:            &S3UserDeleteRequest{UserID: "1"},
		},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3UserDeleteRequest{UserID: "1"},
			}},
	)

}

func TestS3UserDeleteRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserDeleteRequest{}, t)
}
