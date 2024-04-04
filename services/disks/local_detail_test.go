package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestLocalDetailRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &LocalDetailRequest{LocalDiskID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(localDetailEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestLocalDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&LocalDetailRequest{}, t)
}

func TestLocalDetailRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id": "disk_id", "name": "name"}}`, http.StatusOK
				},
				Req:      &LocalDetailRequest{LocalDiskID: "disk_id"},
				Expected: &LocalDiskDetailResponse{Result: LocalDisk{ID: "disk_id", Name: "name"}},
				Actual:   &LocalDiskDetailResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &LocalDetailRequest{LocalDiskID: "disk_id"},
			},
		},
	)
}
