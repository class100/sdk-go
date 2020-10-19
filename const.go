package class100

import (
	`github.com/storezhang/gox`
)

const (
	// MaxRetryTimes 最大重试次数
	MaxRetryTimes int = 100
	// DefaultRetryTimes 默认重试次数
	DefaultRetryTimes int = 6
)

const (
	// ApiVersionDefault 默认版本
	ApiVersionDefault ApiVersion = "default"
	// ApiVersionV1 API版本
	ApiVersionV1 ApiVersion = "v1"
)

const (
	// ErrorCodeValidate 数据验证错误
	ErrorCodeValidate gox.ErrorCode = 401
)
