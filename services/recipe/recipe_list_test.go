package recipe

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	intTesting "github.com/clo-ru/cloapi-go-client/internal/testing"
	"github.com/clo-ru/cloapi-go-client/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRecipeListRequest_BuildRequest(t *testing.T) {
	projectID := "project_id"
	req := &RecipeListRequest{ProjectID: projectID}
	intTesting.BuildTest(req, http.MethodGet, fmt.Sprintf(recipeListEndpoint, mocks.MockUrl, projectID), nil, t)
}

func TestRecipeListRequest_Make(t *testing.T) {
	intTesting.TestDoRequestCases(
		t,
		[]intTesting.DoTestCase{
			{
				Name: "Success",
				BodyStringFunc: func() (string, int) {
					return `{"count": 2,"result": [{"name":"first_item_name","min_disk": 2}, {"name":"second_item_name","suitable_images":["1"]}]}`,
						http.StatusOK
				},
				Req: &RecipeListRequest{},
				Expected: &RecipeListResponse{
					Count: 2,
					Result: []Recipe{
						{
							Name:    "first_item_name",
							MinDisk: 2,
						},
						{
							Name:           "second_item_name",
							SuitableImages: []string{"1"},
						},
					},
				},
				Actual: &RecipeListResponse{},
			},
			{
				Name:       "Error",
				ShouldFail: true,
				BodyStringFunc: func() (string, int) {
					return "", http.StatusInternalServerError
				},
				Req: &RecipeListRequest{},
			},
		},
	)
}

func TestRecipeListRequest_MakeRetry(t *testing.T) {
	intTesting.ConcurrentRetryTest(&RecipeListRequest{}, t)
}

func TestRecipeList_Filtering(t *testing.T) {
	intTesting.FilterTest(&RecipeListRequest{}, t)
}

func TestRecipeListPaginator_lastPage(t *testing.T) {
	httpCli := mocks.RequestDebugClient{}
	cfg := clo.Config{AuthKey: mocks.MockAuthKey, BaseUrl: mocks.MockUrl}
	cli, err := clo.NewDefaultClientFromConfig(cfg)
	if err != nil {
		assert.NoErrorf(t, err, "Client created with error")
	}
	cli.HttpClient = &httpCli
	mocks.BodyStringFunc = func() (string, int) {
		return `{"count": 2,"result": [{"name":"first_item_name","min_disk": 2},{"name":"second_item_name","suitable_images":["1"]}]}`,
			http.StatusOK
	}
	req := &RecipeListRequest{}
	res := &RecipeListResponse{}
	pg := clo.NewPaginator(cli, req, 3, 3)

	assert.Equal(t, false, pg.LastPage())

	err = pg.NextPage(context.Background(), res)
	assert.Nil(t, err)
	assert.Equal(t, true, pg.LastPage())

	err = pg.NextPage(context.Background(), res)
	assert.Equal(t, "no more pages", err.Error())
}
