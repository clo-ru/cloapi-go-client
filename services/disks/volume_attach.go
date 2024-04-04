package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	volumeAttachEndpoint = "%s/v2/volumes/%s/attach"
)

type VolumeAttachRequest struct {
	clo.Request
	VolumeID string
	Body     VolumeAttachBody
}

type VolumeAttachBody struct {
	ServerID string `json:"server_id"`
}

func (r *VolumeAttachRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *VolumeAttachRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(volumeAttachEndpoint, baseUrl, r.VolumeID), authToken, body)
}
