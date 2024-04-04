package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	s3KeysResetEndpoint = "%s/v2/s3/users/%s/credentials"
)

type S3KeysResetRequest struct {
	clo.Request
	UserID string
}

func (r *S3KeysResetRequest) Do(ctx context.Context, cli *clo.ApiClient) (*S3KeysResetResponse, error) {
	res := &S3KeysResetResponse{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *S3KeysResetRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(s3KeysResetEndpoint, baseUrl, r.UserID), authToken, nil)
}
