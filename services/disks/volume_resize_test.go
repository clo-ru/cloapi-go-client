package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestVolumeExtendRequest_BuildRequest(t *testing.T) {
	b := VolumeResizeBody{NewSize: 100}
	dID := "volume_id"
	req := &VolumeResizeRequest{Body: b, VolumeID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(volumeResizeEndpoint, mocks.MockUrl, dID), b, t)
}

func TestVolumeExtendRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &VolumeResizeRequest{VolumeID: "volume_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &VolumeResizeRequest{VolumeID: "volume_id"},
			},
		},
	)
}

func TestVolumeExtendRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&VolumeResizeRequest{}, t)
}
