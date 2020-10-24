package nuwa

import (
	`encoding/json`

	`github.com/class100/sdk-go`
)

// Request 女娲打包请求
type Request struct {
	class100.Request

	// Package 打包参数
	Package Package `json:"package" validate:"required,structonly"`
}

func (r Request) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "    ")

	return string(jsonBytes)
}
