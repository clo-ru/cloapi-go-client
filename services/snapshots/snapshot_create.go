package snapshots

import (
	"context"
	"fmt"
	"github.com/clo-ru/cloapi-go-client/clo"
	tools "github.com/clo-ru/cloapi-go-client/clo/request_tools"
	"net/http"
)

const (
	snapCreateEndpoint = "%s/v2/servers/%s/snapshot"
)

type SnapshotCreateRequest struct {
	clo.Request
	ServerID string
	Body     SnapshotCreateBody
}

type SnapshotCreateBody struct {
	Name string `json:"name"`
}

func (r *SnapshotCreateRequest) Do(ctx context.Context, cli *clo.ApiClient) (*clo.ResponseCreated, error) {
	res := &clo.ResponseCreated{}
	return res, cli.DoRequest(ctx, r, res)
}

func (r *SnapshotCreateRequest) Build(ctx context.Context, baseUrl string, authToken string) (*http.Request, error) {
	body, err := tools.StructToReader(r.Body)
	if err != nil {
		return nil, err
	}
	return r.BuildRaw(ctx, http.MethodPost, fmt.Sprintf(snapCreateEndpoint, baseUrl, r.ServerID), authToken, body)
}
