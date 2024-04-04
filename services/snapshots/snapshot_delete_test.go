package snapshots

import (
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo/mocks"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"net/http"
	"testing"
)

func TestSnapshotDeleteRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &SnapshotDeleteRequest{SnapshotID: ID}
	intTesting.BuildTest(req, http.MethodDelete, fmt.Sprintf(snapDeleteEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestSnapshotDeleteRequest_Make(t *testing.T) {

	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name:           "Success",
				BodyStringFunc: func() (string, int) { return "1", http.StatusOK },
				Req:            &SnapshotDeleteRequest{SnapshotID: "id"},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				CheckError:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &SnapshotDeleteRequest{SnapshotID: "id"},
			},
		},
	)
}

func TestSnapshotDeleteRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&SnapshotDeleteRequest{SnapshotID: "id"}, t)
}
