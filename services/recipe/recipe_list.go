package recipe

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	recipeListEndpoint = "%s/v2/projects/%s/recipes"
)

type RecipeListRequest struct {
	clo.Request
	ProjectID string
}

func (r *RecipeListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*RecipeListResponse, error) {
	resp := &RecipeListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *RecipeListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(recipeListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *RecipeListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *RecipeListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
