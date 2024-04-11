package request_tools

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func IsError(statusCode int) bool {
	switch statusCode {
	case http.StatusOK, http.StatusAccepted, http.StatusCreated, http.StatusNoContent:
		return false
	default:
		return true
	}
}

func StructToReader(obj any) (io.Reader, error) {
	bd := new(bytes.Buffer)
	err := json.NewEncoder(bd).Encode(obj)
	if err != nil {
		return nil, err
	}
	return bd, err
}
