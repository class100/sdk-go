package nuwa

import (
	`github.com/storezhang/gox`
)

// ErrorNotSupportPackage 不支持的打包类型
var ErrorNotSupportPackage = &gox.CodeError{ErrorCode: 101, Msg: "不支持的打包类型"}
