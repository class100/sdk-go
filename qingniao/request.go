package qingniao

import (
	"encoding/json"

	"github.com/class100/sdk-go"
)

// Request 青鸟请求
type Request struct {
	class100.Request

	// Notify 请求
	Notify Notify `json:"notify" validate:"required,structonly"`
}

func (r Request) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "    ")

	return string(jsonBytes)
}
