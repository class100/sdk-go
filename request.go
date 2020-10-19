package class100

import (
	`encoding/json`
)

// Request 请求
type Request struct {
	// Channel 请求通道
	Channel Channel `json:"channel" validate:"required,oneof=dev test prod"`
}

func (r Request) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "    ")

	return string(jsonBytes)
}
