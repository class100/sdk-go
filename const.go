package class100

import (
	"github.com/storezhang/gox"
)

const (
	// TagPackageWindows 用于Windows打包
	TagPackageWindows string = "package-windows"
	// TagPackageMac 用于Mac打包
	TagPackageMac string = "package-mac"
	// TagPackageAndroid 用于Android打包
	TagPackageAndroid string = "package-android"
	// TagPackageIOS 用于iOS打包
	TagPackageIOS string = "package-ios"

	// TagNotify 用于通知
	TagNotify string = "notify"

	// MaxRetryTimes 最大重试次数
	MaxRetryTimes int = 100
	// DefaultRetryTimes 默认重试次数
	DefaultRetryTimes int = 6
)

const (
	// ChannelDev 开发
	ChannelDev Channel = "dev"
	// ChannelTest 测试
	ChannelTest Channel = "test"
	// ChannelProd 生产
	ChannelProd Channel = "prod"
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
