package nuwa

import (
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
		Payload map[string]string `json:"payload"`
	}

	// NotifyRequest 回调请求
	NotifyRequest struct {
		// Status 打包状态
		Status PackageStatus `json:"status" validate:"required"`
		// SrcFile 源文件
		SrcFile transfer.File `json:"srcFile" validate:"required"`
		// DestFile 打包后的文件
		DestFile transfer.File `json:"destFile" validate:"required"`
		// Payload 透传数据
		Payload map[string]string `json:"payload"`
	}
)
