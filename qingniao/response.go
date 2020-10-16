package qingniao

import (
	`encoding/json`
)

// Response 响应
type Response struct {
	// 编号
	Id string `json:"id"`
	// 关键信息
	Key string `json:"key"`
}

func (r Response) String() string {
	jsonBytes, _ := json.MarshalIndent(r, "", "    ")

	return string(jsonBytes)
}
