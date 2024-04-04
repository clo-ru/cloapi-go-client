package clo

import (
	"encoding/json"
	"io"
)

type ResponseInterface interface{}
type ListResponseInterface interface {
	GetCount() int
}

type Response[T any] struct {
	Result T `json:"result,omitempty"`
}

type ListResponse[T any] struct {
	Result []T `json:"result,omitempty"`
	Count  int `json:"count,omitempty"`
}

func (lr *ListResponse[T]) GetCount() int {
	return lr.Count
}

type IdResult struct {
	ID string `json:"id,omitempty"`
}

type ResponseCreated = Response[IdResult]

func UnmarshallJsonResponse(body io.Reader, dst any) error {
	return json.NewDecoder(body).Decode(dst)
}
