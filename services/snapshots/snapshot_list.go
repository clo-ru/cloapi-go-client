package snapshots

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	snapListEndpoint = "%s/v2/projects/%s/snapshots"
)

type SnapshotListRequest struct {
	clo.Request
	ProjectID string
}

func (r *SnapshotListRequest) Do(ctx context.Context, cli *clo.ApiClient) (*SnapshotDetailResponse, error) {
	resp := &SnapshotDetailResponse{}
	return resp, cli.DoRequest(ctx, r, resp)
}

func (r *SnapshotListRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(snapListEndpoint, baseUrl, r.ProjectID), authToken, nil)
}

func (r *SnapshotListRequest) OrderBy(of string)              { r.WithQueryParams(clo.QueryParam{"order": {of}}) }
func (r *SnapshotListRequest) FilterBy(ff clo.FilteringField) { clo.AddFilterToRequest(r, ff) }
