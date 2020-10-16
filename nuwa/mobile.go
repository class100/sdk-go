package nuwa

import (
	`encoding/json`

	`github.com/storezhang/transfer`
)

// Splash 闪屏图片
type Splash struct {
	// 文件名，使用相对路径
	FileName string `json:"fileName"`
	// 文件，真实的图片
	File transfer.File `json:"File"`
}

func (s Splash) String() string {
	jsonBytes, _ := json.MarshalIndent(s, "", "    ")

	return string(jsonBytes)
}
