package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestVolumeListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &VolumeListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(volumeListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestVolumeListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2, "result": [{"id": "first_item_id", "name": "first_item_name"},{"id": "second_item_id", "name": "second_item_name"}]}`, http.StatusOK
				},
				Req: &VolumeListRequest{ProjectID: "project_id"},
				Expected: &VolumeListResponse{
					Count: 2,
					Result: []Volume{
						{ID: "first_item_id", Name: "first_item_name"},
						{ID: "second_item_id", Name: "second_item_name"},
					},
				},
				Actual: &VolumeListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &VolumeListRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestVolumeListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&VolumeListRequest{}, t)
}

func TestVolumeList_Filtering(t *testing.T) {
	intTesting.FilterTest(&VolumeListRequest{}, t)
}
