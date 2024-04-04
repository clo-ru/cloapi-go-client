package storage

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo/mocks"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"net/http"
	"testing"
)

func TestS3UserSuspendRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &S3UserSuspendRequest{UserID: ID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(s3UserSuspendEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestS3UserSuspendRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &S3UserSuspendRequest{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3UserSuspendRequest{},
			},
		},
	)
}

func TestS3UserSuspendRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserSuspendRequest{}, t)
}
