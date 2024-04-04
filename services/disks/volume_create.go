package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	volumeCreateEndpoint = "%s/v2/projects/%s/volumes"
)

type VolumeCreateRequest struct {
	clo.Request
	ProjectID string
	Body      VolumeCreateBody
}

type VolumeCreateBody struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Autorename bool   `json:"autorename"`
}

func (r *VolumeCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *VolumeCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(volumeCreateEndpoint, baseUrl, r.ProjectID), authToken, body)
}
