package snapshots

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	snapRestoreEndpoint = "%s/v2/snapshots/%s/restore"
)

type SnapshotRestoreRequest struct {
	clo.Request
	SnapshotID string
	Body       SnapshotRestoreBody
}

type SnapshotRestoreBody struct {
	Name string `json:"name"`
}

func (r *SnapshotRestoreRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *SnapshotRestoreRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(snapRestoreEndpoint, baseUrl, r.SnapshotID), authToken, body)
}
