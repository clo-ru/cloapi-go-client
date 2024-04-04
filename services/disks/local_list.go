package disks

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	localListEndpoint = "%s/v2/projects/%s/local-disks"
)

type LocalListRequest struct {
	clo.Request
	ProjectID string
}

func (r *LocalListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*LocalDiskListResponse, error) {
	resp := &LocalDiskListResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *LocalListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(localListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *LocalListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *LocalListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
