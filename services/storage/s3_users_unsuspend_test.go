package storage

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3UserUnsuspendRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &S3UserUnsuspendRequest{UserID: ID}

	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(s3UserUnsuspendEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestS3UserUnsuspendRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &S3UserUnsuspendRequest{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3UserUnsuspendRequest{},
			}},
	)
}

func TestS3UserUnsuspendRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserUnsuspendRequest{}, t)
}
