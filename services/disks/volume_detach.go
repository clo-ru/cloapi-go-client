package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	volumeDetachEndpoint = "%s/v2/volumes/%s/detach"
)

type VolumeDetachRequest struct {
	clo.Request
	VolumeID string
	Body     VolumeDetachBody
}

type VolumeDetachBody struct {
	Force bool `json:"force"`
}

func (r *VolumeDetachRequest) Do(ctx context.Context, cli *clo.ApiClient) error {
	return cli.DoRequest(ctx, r, nil)
}

func (r *VolumeDetachRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(volumeDetachEndpoint, baseUrl, r.VolumeID), authToken, body)
}
