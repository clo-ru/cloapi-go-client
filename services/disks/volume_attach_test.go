package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestVolumeAttachRequest_BuildRequest(t *testing.T) {
	b := VolumeAttachBody{ServerID: "id"}
	dID := "volume_id"
	req := &VolumeAttachRequest{Body: b, VolumeID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(volumeAttachEndpoint, mocks.MockUrl, dID), b, t)
}

func TestVolumeAttachRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return "", http.StatusOK
				},
				Req: &VolumeAttachRequest{VolumeID: "disk_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &VolumeAttachRequest{VolumeID: "disk_id"},
			},
		},
	)
}

func TestVolumeAttachRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&VolumeAttachRequest{}, t)
}
