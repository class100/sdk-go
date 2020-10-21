package class100

import (
	`github.com/storezhang/gox`
)

const (
	// MaxRetryTimes 最大重试次数
	MaxRetryTimes int = 10
	// DefaultRetryTimes 默认重试次数
	DefaultRetryTimes int = 3
)

const (
	// ApiVersionDefault 默认版本
	ApiVersionDefault ApiVersion = "default"
	// ApiVersionV1 V1版本
	ApiVersionV1 ApiVersion = "v1"
	// ApiVersionV2 V2版本
	ApiVersionV2 ApiVersion = "v2"
)

const (
	// ErrorCodeValidate 数据验证错误
	ErrorCodeValidate gox.ErrorCode = 401
)
