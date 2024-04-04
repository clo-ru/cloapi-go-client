package request_tools

import (
	"fmt"
)

type DefaultError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

func (de DefaultError) Error() string {
	return fmt.Sprintf("error on api request Code: %d Message: %s Path: %s", de.Code, de.Message, de.Path)
}
