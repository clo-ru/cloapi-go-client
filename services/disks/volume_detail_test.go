package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestVolumeDetailRequest_BuildRequest(t *testing.T) {
	dID := "volume_id"
	req := &VolumeDetailRequest{VolumeID: dID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(volumeDetailEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestVolumeDetailRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"result":{"id":"disk_id","name":"some_name","device":"sda","undetachable":true,"attached_to_server":{"id":"server", "device": "sda"}}}`, http.StatusOK
				},
				Req: &VolumeDetailRequest{VolumeID: "id"},
				Expected: &VolumeDetailResponse{
					Result: Volume{
						ID:           "disk_id",
						Name:         "some_name",
						Undetachable: true,
						Attachment:   &DiskAttachment{"server", "sda"},
					},
				},
				Actual: &VolumeDetailResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				CheckError: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &VolumeDetailRequest{VolumeID: "id"},
			},
		},
	)
}

func TestVolumeDetailRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&VolumeDetailRequest{}, t)
}
