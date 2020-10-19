package nuwa

import (
	`encoding/json`

	`github.com/storezhang/replace`
	`github.com/storezhang/transfer`

	`github.com/class100/sdk-go`
)

// Windows Windows打包信息
type Windows struct {
	// 安装过程中显示的应用程序名称
	ProductName string `json:"productName" validate:"omitempty"`
	// 安装过程中显示的版本号
	ProductVersion string `json:"productVersion" validate:"omitempty"`
	// 应用程序出版人
	ProductPublisher string `json:"productPublisher" validate:"omitempty"`
	// 应用程序网站
	ProductWebSite string `json:"productWebSite" validate:"omitempty,url"`
	// 安装目录下Exe名称
	RunFileName string `json:"runFileName" validate:"omitempty"`
	// 安装完成后快捷方式的名称
	ShortcutName string `json:"shortcutName" validate:"omitempty"`
	// 安装目录文件夹名
	InstallDirName string `json:"installDirName" validate:"omitempty"`
	// 安装图标
	InstallIcon transfer.File `json:"installIcon" validate:"omitempty,structonly"`
	// 卸载图标
	UninstallIcon transfer.File `json:"uninstallIcon" validate:"omitempty,structonly"`
	// 卸载时的提示语句
	UninstallMessage string `json:"uninstallMessage" validate:"omitempty"`
	// 卸载完成是的提示语句
	UninstallFinishMessage string `json:"uninstallFinishMessage" validate:"omitempty"`
}

// NewSimpleWindowsPackage 创建一个简单的Windows打包
func NewSimpleWindowsPackage(
	windows Windows,
	srcFile transfer.File, destFile transfer.File,
	notify Notify,
	payload interface{},
	replaces ...replace.Replace,
) (pkg *Package, err error) {
	return NewWindowsPackage(windows, class100.DefaultRetryTimes, srcFile, destFile, notify, payload, replaces...)
}

// NewWindowsPackage 创建一个Windows打包
func NewWindowsPackage(
	windows Windows,
	maxRetry int,
	srcFile transfer.File, destFile transfer.File,
	notify Notify,
	payload interface{},
	replaces ...replace.Replace,
) (pkg *Package, err error) {
	return NewPackage(PackageTypeWindows, maxRetry, srcFile, destFile, notify, windows, payload, replaces...)
}

func (w Windows) String() string {
	jsonBytes, _ := json.MarshalIndent(w, "", "    ")

	return string(jsonBytes)
}
