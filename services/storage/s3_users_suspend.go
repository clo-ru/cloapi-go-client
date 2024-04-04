package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	s3UserSuspendEndpoint = "%s/v2/s3/users/%s/suspend"
)

type S3UserSuspendRequest struct {
	clo.Request
	UserID string
}

func (r *S3UserSuspendRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *S3UserSuspendRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(s3UserSuspendEndpoint, baseUrl, r.UserID), authToken, nil)
}
