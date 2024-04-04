package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	s3UserUnsuspendEndpoint = "%s/v2/s3/users/%s/unsuspend"
)

type S3UserUnsuspendRequest struct {
	clo.Request
	UserID string
}

func (r *S3UserUnsuspendRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *S3UserUnsuspendRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(s3UserUnsuspendEndpoint, baseUrl, r.UserID), authToken, nil)
}
