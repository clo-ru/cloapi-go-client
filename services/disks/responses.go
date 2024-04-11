package disks

import "github.com/clo-ru/cloapi-go-client/clo"

type LocalDisk struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Status     string          `json:"status"`
	Size       int             `json:"size"`
	FileSystem string          `json:"file_system"`
	CreatedIn  string          `json:"created_in"`
	Bootable   string          `json:"bootable"`
	Attachment *DiskAttachment `json:"attached_to_server"`
}

type Volume struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Status       string          `json:"status"`
	Size         int             `json:"size"`
	FileSystem   string          `json:"file_system"`
	Description  string          `json:"description"`
	CreatedIn    string          `json:"created_in"`
	Bootable     bool            `json:"bootable"`
	Undetachable bool            `json:"undetachable"`
	Image        *VolumeImage    `json:"image"`
	Recipe       *VolumeRecipe   `json:"recipe"`
	Attachment   *DiskAttachment `json:"attached_to_server"`
	Snapshots    []string        `json:"snapshots"`
}

type VolumeImage struct {
	MinDisk         int                    `json:"min_disk"`
	MinRam          int                    `json:"min_ram"`
	OperationSystem *VolumeOperationSystem `json:"operation_system"`
}

type VolumeRecipe struct {
	Name     string `json:"name"`
	MinDisk  int    `json:"min_disk"`
	MinRam   int    `json:"min_ram"`
	MinVcpus int    `json:"min_vcpus"`
}

type VolumeOperationSystem struct {
	Distribution string `json:"distribution"`
	OsFamily     string `json:"os_family"`
	Version      string `json:"version"`
}

type DiskAttachment struct {
	ID     string `json:"ID"`
	Device string `json:"Device"`
}

type LocalDiskDetailResponse = clo.Response[LocalDisk]
type LocalDiskListResponse = clo.ListResponse[LocalDisk]
type VolumeDetailResponse = clo.Response[Volume]
type VolumeListResponse = clo.ListResponse[Volume]
