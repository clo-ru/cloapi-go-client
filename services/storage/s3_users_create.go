package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const s3UserCreateEndpoint = "%s/v2/projects/%s/s3/users"

type S3UserCreateRequest struct {
	clo.Request
	ProjectID string
	Body      S3UserCreateBody
}

type S3UserCreateBody struct {
	Name          string            `json:"name"`
	CanonicalName string            `json:"canonical_name"`
	DefaultBucket bool              `json:"default_bucket"`
	MaxBuckets    int               `json:"max_buckets"`
	BucketQuota   CreateQuotaParams `json:"bucket_quota,omitempty"`
	UserQuota     CreateQuotaParams `json:"user_quota"`
}

type CreateQuotaParams struct {
	MaxObjects int `json:"max_objects,omitempty"`
	MaxSize    int `json:"max_size"`
}

func (r *S3UserCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	resp := &clo.ResponseCreated{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *S3UserCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(s3UserCreateEndpoint, baseUrl, r.ProjectID), authToken, body)
}
