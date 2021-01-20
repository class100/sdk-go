package class100

import (
	"encoding/json"

	"github.com/class100/core"
)

// Request 请求
type Request struct {
	// Environment 环境类型
	Environment core.Environment `json:"environment" validate:"required,oneof=dev test prod local qa"`
}

func (r Request) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "    ")

	return string(jsonBytes)
}
