package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	s3KeysGetEndpoint = "%s/v2/s3/users/%s/credentials"
)

type S3KeysGetRequest struct {
	clo.Request
	UserID string
}

func (r *S3KeysGetRequest) Do(ctx context.Context, cli *clo.ApiClient) (*S3KeysGetResponse, error) {
	resp := &S3KeysGetResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *S3KeysGetRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(s3KeysGetEndpoint, baseUrl, r.UserID), authToken, nil)
}
