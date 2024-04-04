package storage

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3UserQuotaPatchRequest_BuildRequest(t *testing.T) {
	b := S3UserQuotaPatchBody{}
	ID := "id"
	req := &S3UserQuotaPatchRequest{UserID: ID, Body: b}
	intTesting.BuildTest(req, http.MethodPut, fmt.Sprintf(s3UserQuotaPatchEndpoint, mocks.MockUrl, ID), b, t)
}

func TestS3UserQuotaPatchRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return "1", http.StatusOK
				},
				Req: &S3UserQuotaPatchRequest{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3UserQuotaPatchRequest{},
			}},
	)
}

func TestS3UserQuotaPatchRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserQuotaPatchRequest{}, t)
}
