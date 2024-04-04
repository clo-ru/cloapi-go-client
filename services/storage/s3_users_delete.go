package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	s3UserDeleteEndpoint = "%s/v2/s3/users/%s"
)

type S3UserDeleteRequest struct {
	clo.Request
	UserID string
}

func (r *S3UserDeleteRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *S3UserDeleteRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodDelete, fmt.Sprintf(s3UserDeleteEndpoint, baseUrl, r.UserID), authToken, nil)
}
