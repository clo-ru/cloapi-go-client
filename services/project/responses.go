package project

import "github.com/clo-ru/cloapi-go-client/clo"

type Project struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	CreatedIn      string `json:"created_in"`
	StoppingReason string `json:"stopping_reason"`
	HasAbuse       bool   `json:"has_abuse"`
}

type Image struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	OperationSystem *OperationSystem `json:"operation_system"`
}

type OperationSystem struct {
	Distribution string `json:"distribution"`
	OsFamily     string `json:"os_family"`
	Version      string `json:"version"`
}

type Limit struct {
	Max     int    `json:"max"`
	Used    int    `json:"used"`
	Name    string `json:"name"`
	Measure string `json:"measure"`
}

type LimitsListResponse = clo.ListResponse[Limit]
type ImageListResponse = clo.ListResponse[Image]
type ProjectListResponse = clo.ListResponse[Project]
