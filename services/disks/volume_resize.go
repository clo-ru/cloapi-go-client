package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	volumeResizeEndpoint = "%s/v2/volumes/%s/extend"
)

type VolumeResizeRequest struct {
	clo.Request
	VolumeID string
	Body     VolumeResizeBody
}

type VolumeResizeBody struct {
	NewSize int `json:"new_size"`
}

func (r *VolumeResizeRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *VolumeResizeRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(volumeResizeEndpoint, baseUrl, r.VolumeID), authToken, body)
}
