package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestLocalListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &LocalListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(localListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestLocalListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2, "result": [{"id": "first_item_id", "name": "first_item_name"},{"id": "second_item_id", "name": "second_item_name"}]}`, http.StatusOK
				},
				Req: &LocalListRequest{
					ProjectID: "project_id",
				},
				Expected: &LocalDiskListResponse{
					Count: 2,
					Result: []LocalDisk{
						{ID: "first_item_id", Name: "first_item_name"},
						{ID: "second_item_id", Name: "second_item_name"},
					},
				},
				Actual: &LocalDiskListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &LocalListRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestLocalListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&LocalListRequest{}, t)
}

func TestLocalList_Filtering(t *testing.T) {
	intTesting.FilterTest(&LocalListRequest{}, t)
}
