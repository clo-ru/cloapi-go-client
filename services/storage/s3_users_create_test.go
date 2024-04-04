package storage

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3UserCreateRequest_BuildRequest(t *testing.T) {
	b := S3UserCreateBody{Name: "m"}
	ID := "id"
	req := &S3UserCreateRequest{ProjectID: ID, Body: b}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(s3UserCreateEndpoint, mocks.MockUrl, ID), b, t)
}

func TestS3UserCreateRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id":"user_id"}}`, http.StatusOK
				},
				Req:      &S3UserCreateRequest{},
				Expected: &clo.ResponseCreated{Result: clo.IdResult{ID: "user_id"}},
				Actual:   &clo.ResponseCreated{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &S3UserCreateRequest{},
			},
		},
	)
}

func TestS3UserCreateRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserCreateRequest{}, t)
}
