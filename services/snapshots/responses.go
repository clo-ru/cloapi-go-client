package snapshots

import "github.com/clo-ru/cloapi-go-client/clo"

type Snapshot struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	CreatedIn    string   `json:"created_in"`
	DeletedIn    string   `json:"deleted_in"`
	Status       string   `json:"status"`
	Size         int      `json:"size"`
	ParentServer string   `json:"parent_server"`
	ChildServers []string `json:"child_servers"`
}

type SnapshotDetailResponse = clo.Response[Snapshot]
type SnapshotListResponse = clo.ListResponse[Snapshot]
