package snapshots

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestSnapshotListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &SnapshotListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(snapListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestSnapshotListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2,"result": [{"name":"first_item_name","parent_server":"server1_id"},{"name":"second_item_name","child_servers":["server2_id"]}]}`, http.StatusOK
				},
				Req: &SnapshotListRequest{},
				Expected: &SnapshotListResponse{
					Count: 2,
					Result: []Snapshot{
						{
							Name:         "first_item_name",
							ParentServer: "server1_id",
						},
						{
							Name:         "second_item_name",
							ChildServers: []string{"server2_id"},
						},
					},
				},
				Actual: &SnapshotListResponse{},
			},
			{
				Name:           "Error",
				ShouldFail:     true,
				BodyStringFunc: func() (string, int) { return "", http.StatusInternalServerError },
				Req:            &SnapshotListRequest{},
			},
		},
	)

}

func TestSnapshotListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&SnapshotListRequest{}, t)
}

func TestServerList_Filtering(t *testing.T) {
	intTesting.FilterTest(&SnapshotListRequest{}, t)
}
