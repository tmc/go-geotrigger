package geotrigger

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%+v", e)
}

type errorResponse struct {
	Error Error `json:"error"`
}
