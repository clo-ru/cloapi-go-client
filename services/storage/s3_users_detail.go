package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	s3UserDetailEndpoint = "%s/v2/s3/users/%s/detail"
)

type S3UserDetailRequest struct {
	clo.Request
	UserID string
}

func (r *S3UserDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*S3UserDetailResponse, error) {
	resp := &S3UserDetailResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *S3UserDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(s3UserDetailEndpoint, baseUrl, r.UserID), authToken, nil)
}
