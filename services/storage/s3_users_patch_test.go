package storage

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3UserPatchRequest_BuildRequest(t *testing.T) {
	b := S3UserPatchBody{Name: "m"}
	ID := "id"
	req := &S3UserPatchRequest{UserID: ID, Body: b}
	intTesting.BuildTest(req, http.MethodPatch, fmt.Sprintf(s3UserPatchEndpoint, mocks.MockUrl, ID), b, t)
}

func TestS3UserPatchRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &S3UserPatchRequest{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &S3UserPatchRequest{},
			},
		},
	)
}

func TestS3UserPatchRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserPatchRequest{}, t)
}
