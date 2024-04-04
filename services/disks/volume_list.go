package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	volumeListEndpoint = "%s/v2/projects/%s/volumes"
)

type VolumeListRequest struct {
	clo.Request
	ProjectID string
}

func (r *VolumeListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*VolumeListResponse, error) {
	resp := &VolumeListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}
func (r *VolumeListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(volumeListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}
func (r *VolumeListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *VolumeListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
