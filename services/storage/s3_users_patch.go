package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	s3UserPatchEndpoint = "%s/v2/s3/users/%s"
)

type S3UserPatchRequest struct {
	clo.Request
	UserID string
	Body   S3UserPatchBody
}

type S3UserPatchBody struct {
	Name string `json:"name"`
}

func (r *S3UserPatchRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *S3UserPatchRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPatch, fmt.Sprintf(s3UserPatchEndpoint, baseUrl, r.UserID), authToken, body)
}
