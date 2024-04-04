package project

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestProjectListRequest_BuildRequest(t *testing.T) {
	req := &ProjectListRequest{}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(projectListEndpoint, mocks.MockUrl), nil, t)
}

func TestProjectListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2,"result": [{"name":"first_item_name"},{"name":"second_item_name"}]}`, http.StatusOK
				},
				Req: &ProjectListRequest{},
				Expected: &ProjectListResponse{
					Count: 2,
					Result: []Project{
						{Name: "first_item_name"},
						{Name: "second_item_name"},
					},
				},
				Actual: &ProjectListResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req:      &ProjectListRequest{},
				Expected: nil,
			},
		},
	)

}

func TestProjectListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&ProjectListRequest{}, t)
}

func TestProjectList_Filtering(t *testing.T) {
	intTesting.FilterTest(&ProjectListRequest{}, t)
}
