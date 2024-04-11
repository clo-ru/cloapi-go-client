package servers

import "github.com/clo-ru/cloapi-go-client/clo"

type Server struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Status         string        `json:"status"`
	CreatedIn      string        `json:"created_in"`
	RescueMode     string        `json:"rescue_mode"`
	Datacenter     string        `json:"datacenter"`
	PrimaryAddress string        `json:"primary_address"`
	Project        string        `json:"project"`
	GuestAgent     bool          `json:"guest_agent"`
	AdminKey       string        `json:"admin_key_status"`
	Flavor         ServerFlavor  `json:"flavor"`
	Image          *ServerImage  `json:"image"`
	Recipe         *ServerRecipe `json:"recipe"`
	DiskData       []ServerDisk  `json:"disk_data"`
	Addresses      []string      `json:"addresses"`
	Snapshots      []string      `json:"snapshots"`
}

type ServerRecipe struct {
	MinDisk  int    `json:"min_disk"`
	MinRam   int    `json:"min_ram"`
	MinVcpus int    `json:"min_vcpus"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}

type ServerImage struct {
	MinDisk         int           `json:"min_disk"`
	MinRam          int           `json:"min_ram"`
	OperationSystem *ServerSystem `json:"operation_system"`
}

type ServerSystem struct {
	Distribution string `json:"distribution"`
	OsFamily     string `json:"os_family"`
	Version      string `json:"version"`
}

type ServerFlavor struct {
	Ram     int    `json:"ram"`
	Vcpus   int    `json:"vcpus"`
	Disk    int    `json:"disk"`
	CpyType string `json:"cpu_type"`
}

type ServerDisk struct {
	ID          string `json:"id"`
	StorageType string `json:"storage_type"`
}

type ServerDetailResponse = clo.Response[Server]
type ServerListResponse = clo.ListResponse[Server]
