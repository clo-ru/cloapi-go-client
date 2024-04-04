package storage

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestS3UserDetailRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &S3UserDetailRequest{UserID: ID}

	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(s3UserDetailEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestS3UserDetailRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return fmt.Sprintf(`{"result":{"id":"user_id","max_buckets":2,"quotas":[{"max_size":1,"max_objects":2,"type":"user"}]}}`),
						http.StatusOK
				},
				Req: &S3UserDetailRequest{},
				Expected: &S3UserDetailResponse{
					Result: S3User{
						ID:         "user_id",
						MaxBuckets: 2,
						Quotas: []QuotaInfo{
							{
								MaxSize:    1,
								MaxObjects: 2,
								Type:       "user",
							},
						},
					},
				},
				Actual: &S3UserDetailResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &S3UserDetailRequest{},
			}},
	)
}

func TestS3UserDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&S3UserDetailRequest{}, t)
}
