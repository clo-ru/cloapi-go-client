package snapshots

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	"net/http"
)

const (
	snapDeleteEndpoint = "%s/v2/snapshots/%s"
)

type SnapshotDeleteRequest struct {
	clo.Request
	SnapshotID string
}

func (r *SnapshotDeleteRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *SnapshotDeleteRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	return r.BuildRaw(ctx, http.MethodDelete, fmt.Sprintf(snapDeleteEndpoint, baseUrl, r.SnapshotID), authToken, nil)
}
