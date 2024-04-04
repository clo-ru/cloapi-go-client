package snapshots

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	snapDetailEndpoint = "%s/v2/snapshots/%s/detail"
)

type SnapshotDetailRequest struct {
	clo.Request
	SnapshotID string
}

func (r *SnapshotDetailRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *SnapshotDetailRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodGet, fmt.Sprintf(snapDetailEndpoint, baseUrl, r.SnapshotID), authToken, nil)
}
