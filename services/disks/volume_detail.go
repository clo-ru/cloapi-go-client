package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	volumeDetailEndpoint = "%s/v2/volumes/%s/detail"
)

type VolumeDetailRequest struct {
	clo.Request
	VolumeID string
}

func (r *VolumeDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*VolumeDetailResponse, error) {
	res := &VolumeDetailResponse{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *VolumeDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(volumeDetailEndpoint, baseUrl, r.VolumeID), authToken, nil)
}
