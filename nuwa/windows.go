package nuwa

import (
	`encoding/json`

	`github.com/storezhang/replace`
	`github.com/storezhang/transfer`

	`github.com/class100/sdk-go`
)

// Windows Windows打包信息
type Windows struct {
	BasePackager

	// ProductId 产品的唯一标识
	ProductId string `json:"productId" validate:"required,start_with_alpha"`
	// ProductName 安装过程中显示的应用程序名称
	ProductName string `json:"productName" validate:"omitempty"`
	// ProductVersion 安装过程中显示的版本号
	ProductVersion string `json:"productVersion" validate:"omitempty"`
	// ProductPublisher 应用程序出版人
	ProductPublisher string `json:"productPublisher" validate:"omitempty"`
	// ProductWebsite 应用程序网站
	ProductWebsite string `json:"productWebsite" validate:"omitempty,url"`
	// RunFilename 安装目录下Exe名称
	RunFilename string `json:"runFilename" validate:"omitempty"`
	// ShortcutName 安装完成后快捷方式的名称
	ShortcutName string `json:"shortcutName" validate:"omitempty"`
	// InstallDirName 安装目录文件夹名
	InstallDirName string `json:"installDirName" validate:"omitempty"`
	// InstallIcon 安装图标
	InstallIcon transfer.File `json:"installIcon" validate:"omitempty,structonly"`
	// UninstallIcon 卸载图标
	UninstallIcon transfer.File `json:"uninstallIcon" validate:"omitempty,structonly"`
	// UninstallMessage 卸载时的提示语句
	UninstallMessage string `json:"uninstallMessage" validate:"omitempty"`
	// UninstallFinishMessage 卸载完成是的提示语句
	UninstallFinishMessage string `json:"uninstallFinishMessage" validate:"omitempty"`
}

// NewSimpleWindowsPackage 创建一个简单的Windows打包
func NewSimpleWindowsPackage(
	windows Windows,
	srcFile transfer.File, destFile transfer.File,
	notify Notify,
	replaces ...replace.Replace,
) *Package {
	return NewWindowsPackage(windows, class100.DefaultRetryTimes, srcFile, destFile, notify, replaces...)
}

// NewWindowsPackage 创建一个Windows打包
func NewWindowsPackage(
	windows Windows,
	maxRetry int,
	srcFile transfer.File, destFile transfer.File,
	notify Notify,
	replaces ...replace.Replace,
) *Package {
	return NewPackage(PackageTypeWindows, maxRetry, srcFile, destFile, notify, windows, replaces...)
}

func (w Windows) String() string {
	jsonBytes, _ := json.MarshalIndent(w, "", "    ")

	return string(jsonBytes)
}
