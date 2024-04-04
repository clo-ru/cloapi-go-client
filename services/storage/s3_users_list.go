package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	s3UserListEndpoint = "%s/v2/projects/%s/s3/users"
)

type S3UserListRequest struct {
	clo.Request
	ProjectID string
}

func (r *S3UserListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*S3UserListResponse, error) {
	resp := &S3UserListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *S3UserListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(s3UserListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *S3UserListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *S3UserListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
