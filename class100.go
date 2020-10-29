package class100

import (
	`encoding/json`
)

type (
	// ApiVersion API版本
	ApiVersion string

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
