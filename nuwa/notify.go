package nuwa

import (
	"github.com/storezhang/transfer"
)

type (
	// Notify 发送给通知服务器的请求
	Notify struct {
		// Status 打包状态
		Status PackageStatus `json:"status" validate:"required"`
		// SrcFile 源文件
		SrcFile transfer.File `json:"srcFile" validate:"required"`
		// DestFile 打包后的文件
		DestFile transfer.File `json:"destFile" validate:"required"`
	}
)
