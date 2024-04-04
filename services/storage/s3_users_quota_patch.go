package storage

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	s3UserQuotaPatchEndpoint = "%s/v2/s3/users/%s/quotas"
)

type S3UserQuotaPatchRequest struct {
	clo.Request
	UserID string
	Body   S3UserQuotaPatchBody
}

type S3UserQuotaPatchBody struct {
	MaxBuckets  int              `json:"max_buckets"`
	UserQuota   PatchQuotaParams `json:"user_quota"`
	BucketQuota PatchQuotaParams `json:"bucket_quota"`
}

type PatchQuotaParams struct {
	MaxObjects int `json:"max_objects"`
	MaxSize    int `json:"max_size"`
}

func (r *S3UserQuotaPatchRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *S3UserQuotaPatchRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPut, fmt.Sprintf(s3UserQuotaPatchEndpoint, baseUrl, r.UserID), authToken, body)
}
