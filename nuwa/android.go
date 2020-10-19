package nuwa

import (
	`encoding/json`

	`github.com/storezhang/replace`
	`github.com/storezhang/transfer`

	`github.com/class100/sdk-go`
)

const (
	// DefaultAppNameKey 默认语言键
	DefaultAppNameKey string = "default"
)

type (
	// AndroidAppSign 安卓签名
	AndroidAppSign struct {
		// KeystoreFile 密钥文件
		KeystoreFile transfer.File `json:"keystoreFile"`
		// StorePass 密码
		StorePass string `json:"storePass"`
		// DigestAlg 加密算法
		DigestAlg string `default:"SHA1" json:"digestAlg"`
		// SigAlg 签名算法
		SigAlg string `default:"SHA1withRSA" json:"sigAlg"`
		// Alias 别名
		Alias string `json:"alias"`
	}

	// Android APK打包信息
	Android struct {
		// Name 应用名称
		Name map[string]string `json:"name"`
		// Package 包名
		Package string `json:"package"`
		// Icon 图标
		Icon transfer.File `json:"icon"`
		// Version 版本号
		Version string `json:"version"`
		// Sign 签名
		Sign AndroidAppSign `json:"sign"`
	}
)

// NewSimpleAndroidPackage 创建一个简单的Android打包
func NewSimpleAndroidPackage(
	android Android,
	srcFile transfer.File, destFile transfer.File,
	notify Notify,
	payload interface{},
	replaces ...replace.Replace,
) (pkg *Package, err error) {
	return NewAndroidPackage(android, class100.DefaultRetryTimes, srcFile, destFile, notify, payload, replaces...)
}

// NewAndroidPackage 创建一个Android打包
func NewAndroidPackage(
	android Android,
	maxRetry int,
	srcFile transfer.File, destFile transfer.File,
	notify Notify,
	payload interface{},
	replaces ...replace.Replace,
) (pkg *Package, err error) {
	return NewPackage(PackageTypeAndroid, maxRetry, srcFile, destFile, notify, android, payload, replaces...)
}

func (a Android) String() string {
	jsonBytes, _ := json.MarshalIndent(a, "", "    ")

	return string(jsonBytes)
}
