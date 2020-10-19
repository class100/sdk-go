package nuwa

import (
	`encoding/json`

	`github.com/storezhang/transfer`
)

type (
	// Notify 通知，通过JWT回调
	Notify struct {
		// Url 回调地址
		Url string `json:"url" validate:"required"`
		// Scheme
		Scheme string `default:"Bearer" json:"scheme" validate:"required"`
		// Token JWT验证授权码
		Token string `json:"token" validate:"token"`
		// Payload 透传数据
		Payload []byte `json:"payload"`
	}

	// NotifyRequest 回调请求数据
	NotifyRequest struct {
		// Status 打包状态
		Status PackageStatus `json:"status"`
		// SrcFile 源文件
		SrcFile transfer.File `json:"srcFile"`
		// DestFile 打包后的文件
		DestFile transfer.File `json:"destFile"`
		// Payload 透传数据
		Payload []byte `json:"payload"`
	}
)

// NewNotify 创建一个新的通知
func NewNotify(url string, scheme string, token string, payload interface{}) (notify Notify, err error) {
	var jsonBytes []byte
	if jsonBytes, err = json.Marshal(payload); nil != err {
		return
	}

	notify = Notify{
		Url:     url,
		Scheme:  scheme,
		Token:   token,
		Payload: jsonBytes,
	}

	return
}
