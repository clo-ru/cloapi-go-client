package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestVolumeDeleteRequest_BuildRequest(t *testing.T) {
	dID := "volume_id"
	req := &VolumeDeleteRequest{VolumeID: dID}
	intTesting.BuildTest(req, http.MethodDelete, fmt.Sprintf(volumeDeleteEndpoint, mocks.MockUrl, dID), nil, t)
}

func TestVolumeDeleteRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return "", http.StatusOK
				},
				Req: &VolumeDeleteRequest{VolumeID: "disk_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &VolumeDeleteRequest{VolumeID: "disk_id"},
			},
		},
	)
}

func TestVolumeDeleteRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&VolumeDeleteRequest{}, t)
}
