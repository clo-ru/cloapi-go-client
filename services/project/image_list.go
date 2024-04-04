package project

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	imageListEndpoint = "%s/v2/projects/%s/images"
)

type ImageListRequest struct {
	clo.Request
	ProjectID string
}

func (r *ImageListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*ImageListResponse, error) {
	resp := &ImageListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *ImageListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(imageListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *ImageListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *ImageListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
