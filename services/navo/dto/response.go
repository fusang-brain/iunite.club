package dto

import (
	"encoding/json"
)

type APIResponse struct {
	Code    int64       `json:"Code"`
	SubCode string      `json:"SubCode"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

func (r *APIResponse) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(b)
}
