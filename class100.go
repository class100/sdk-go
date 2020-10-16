package class100

import (
	`encoding/json`
)

type (
	// Channel 通道
	Channel string

	// ApiVersion API版本
	ApiVersion string

	// Request 请求
	Request struct {
		// Channel 请求通道
		Channel Channel `json:"channel" validate:"required,oneof=dev test prod"`
	}

	// Client 客户端
	Client struct {
		// 授权
		AccessKey string `json:"accessKey"`
		SecretKey string `json:"secretKey"`
	}
)

func (c Client) String() string {
	jsonBytes, _ := json.MarshalIndent(c, "", "    ")

	return string(jsonBytes)
}

func (r Request) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "    ")

	return string(jsonBytes)
}
