package qingniao

import (
	`github.com/storezhang/gox`
)

// ErrorNotSupportNotify 不支持的通知类型
var ErrorNotSupportNotify = &gox.CodeError{ErrorCode: 102, Message: "不支持的通知类型"}
