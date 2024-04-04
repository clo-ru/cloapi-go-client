package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	volumeDeleteEndpoint = "%s/v2/volumes/%s"
)

type VolumeDeleteRequest struct {
	clo.Request
	VolumeID string
}

func (r *VolumeDeleteRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *VolumeDeleteRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodDelete, fmt.Sprintf(volumeDeleteEndpoint, baseUrl, r.VolumeID), authToken, nil)
}
