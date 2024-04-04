package storage

import (
	"github.com/clo-ru/cloapi-go-client/clo"
)

type S3Key struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type S3User struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	CanonicalName string      `json:"canonical_name"`
	Status        string      `json:"status"`
	Tenant        string      `json:"tenant"`
	MaxBuckets    int         `json:"max_buckets"`
	Quotas        []QuotaInfo `json:"quotas"`
}

type QuotaInfo struct {
	MaxSize    int    `json:"max_size"`
	MaxObjects int    `json:"max_objects"`
	Type       string `json:"type"`
}

type S3KeysResetResponse = clo.Response[S3Key]
type S3KeysGetResponse = clo.ListResponse[S3Key]
type S3UserDetailResponse = clo.Response[S3User]
type S3UserListResponse = clo.ListResponse[S3User]
