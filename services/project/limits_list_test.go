package project

import (
	"fmt"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"net/http"
	"testing"
)

func TestLimitsListRequest_BuildRequest(t *testing.T) {
	ID := "id"
	req := &LimitsListRequest{ProjectID: ID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(limitsListEndpoint, mocks.MockUrl, ID), nil, t)
}

func TestLimitsListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2,"result": [{"name":"first_item_name","max":2},{"name":"second_item_name","used":3}]}`,
						http.StatusOK
				},
				Req: &LimitsListRequest{ProjectID: "project_id"},
				Expected: &LimitsListResponse{
					Count:  2,
					Result: []Limit{{Max: 2, Name: "first_item_name"}, {Used: 3, Name: "second_item_name"}},
				},
				Actual: &LimitsListResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &LimitsListRequest{ProjectID: "project_id"},
			},
		},
	)
}

func TestLimitsListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&LimitsListRequest{}, t)
}

func TestLimitsListRequest_Filtering(t *testing.T) {
	intTesting.FilterTest(&LimitsListRequest{}, t)
}
