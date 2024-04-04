package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	localDetailEndpoint = "%s/v2/local-disks/%s/detail"
)

type LocalDetailRequest struct {
	clo.Request
	LocalDiskID string
}

func (r *LocalDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*LocalDiskDetailResponse, error) {
	res := &LocalDiskDetailResponse{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *LocalDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(localDetailEndpoint, baseUrl, r.LocalDiskID), authToken, nil)
}
