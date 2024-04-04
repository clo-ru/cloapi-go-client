package disks

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestVolumeDetachRequest_BuildRequest(t *testing.T) {
	b := VolumeDetachBody{Force: true}
	dID := "volume_id"
	req := &VolumeDetachRequest{Body: b, VolumeID: dID}
	intTesting.BuildTest(req, http.MethodPost, fmt.Sprintf(volumeDetachEndpoint, mocks.MockUrl, dID), b, t)
}

func TestVolumeDetachRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "", http.StatusOK },
				Req:            &VolumeDetachRequest{VolumeID: "volume_id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &VolumeDetachRequest{VolumeID: "volume_id"},
			},
		},
	)
}

func TestVolumeDetachRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&VolumeDetachRequest{}, t)
}
