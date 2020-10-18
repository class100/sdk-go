package nuwa

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/storezhang/gox"
	"github.com/storezhang/replace"
	"github.com/storezhang/transfer"

	class100 "github.com/class100/sdk-go"
)

const (
	// 打包类型
	// PackageTypeWindows Windows打包
	PackageTypeWindows PackageType = "windows"
	// PackageTypeMac Mac打包
	PackageTypeMac PackageType = "mac"
	// PackageTypeAndroid 安卓打包
	PackageTypeAndroid PackageType = "android"
	// PackageTypeIOS iOS打包
	PackageTypeIOS PackageType = "ios"
)

type (
	// PackageType 打包类型
	PackageType string

	// Package 打包请求
	Package struct {
		// Type 打包类型
		Type PackageType `json:"type" validate:"required,oneof=windows mac android"`
		// MaxRetry 最大重试次数
		MaxRetry int `default:"3" json:"maxRetry" validate:"omitempty,min=1,max=100"`
		// SrcFile 源文件
		SrcFile transfer.File `json:"srcFile" validate:"required"`
		// DestFile 打包后的文件
		DestFile transfer.File `json:"destFile" validate:"required"`
		// Replaces 文件替换
		Replaces []replace.Replace `json:"replaces" validate:"omitempty,dive"`
		// Notify 通知
		Notify Notify `json:"notify" validate:"omitempty,structonly"`
		// Package 打包数据
		Package interface{} `json:"package" validate:"required"`
		// Payload 透传数据，在Notify时原样提交
		Payload []byte `json:"payload"`
	}
)

func (p *Package) UnmarshalJSON(data []byte) (err error) {
	type cloneType Package

	rawMsg := json.RawMessage{}
	p.Package = &rawMsg

	if err = json.Unmarshal(data, (*cloneType)(p)); err != nil {
		return
	}

	switch p.Type {
	case PackageTypeWindows:
		windows := Windows{}
		if err = json.Unmarshal(rawMsg, &p); err != nil {
			return
		}
		p.Package = windows
	case PackageTypeAndroid:
		android := Android{}
		if err = json.Unmarshal(rawMsg, &p); err != nil {
			return
		}
		p.Package = android
	default:
		err = ErrorNotSupportPackage
	}

	return
}

func (pt PackageType) srcFileExt() string {
	ext := "zip"

	switch pt {
	case PackageTypeWindows:
		ext = "zip"
	case PackageTypeMac:
		ext = "dmg"
	case PackageTypeAndroid:
		ext = "apk"
	case PackageTypeIOS:
		ext = "ipa"
	}

	return ext
}

func (pt PackageType) destFileExt() string {
	ext := "zip"

	switch pt {
	case PackageTypeWindows:
		ext = "exe"
	case PackageTypeMac:
		ext = "dmg"
	case PackageTypeAndroid:
		ext = "apk"
	case PackageTypeIOS:
		ext = "ipa"
	}

	return ext
}

func (p *Package) srcFileName(rootPath string) (srcFileName string) {
	var name string

	switch p.Type {
	case PackageTypeWindows:
		name = p.Package.(Windows).ProductName
	case PackageTypeAndroid:
		name = p.Package.(Android).Name[DefaultAppNameKey]
	}

	srcFileName = filepath.Join(rootPath, gox.GetFileNameWithExt(
		fmt.Sprintf("i-%s-%s", name, strconv.FormatInt(time.Now().UnixNano(), 10)),
		p.Type.srcFileExt(),
	))

	return
}

func (p *Package) destFileName(rootPath string) (destFileName string) {
	var name string

	switch p.Type {
	case PackageTypeWindows:
		name = p.Package.(Windows).ProductName
	case PackageTypeAndroid:
		name = p.Package.(Android).Name[DefaultAppNameKey]
	}

	destFileName = filepath.Join(rootPath, gox.GetFileNameWithExt(
		fmt.Sprintf("o-%s-%s", name, strconv.FormatInt(time.Now().UnixNano(), 10)),
		p.Type.destFileExt(),
	))

	return
}

func (p *Package) packageDir(srcFileName string) (packageDir string) {
	packageDir = gox.GetFileDir(srcFileName)

	switch p.Type {
	case PackageTypeWindows:
		packageDir = fmt.Sprintf("%s-windows", packageDir)
	case PackageTypeAndroid:
		packageDir = fmt.Sprintf("%s-android", packageDir)
	}

	return
}

func (p *Package) Build(rootPath string, packager Packager) (err error) {
	// 验证基本参数
	if err = class100.Validate.Struct(p); nil != err {
		err = gox.NewCodeError(class100.ErrorCodeValidate, "数据验证错误", err.(validator.ValidationErrors))

		return
	}

	srcFileName := p.srcFileName(rootPath)
	// 删除源文件，以免影响下一次打包
	defer func() {
		err = os.RemoveAll(srcFileName)
	}()
	outputFileName := p.destFileName(rootPath)
	packageDir := p.packageDir(srcFileName)
	// 删除打包目录，以免影响下一次打包
	defer func() {
		err = os.RemoveAll(packageDir)
	}()

	// 下载源文件
	if err = p.SrcFile.Download(srcFileName, false); err != nil {
		return
	}
	// 准备
	if err = packager.Decode(srcFileName, packageDir); nil != err {
		return
	}
	// 处理文件替换逻辑
	if err = p.replace(packageDir); nil != err {
		return
	}
	// 处理应用包修改逻辑
	if err = packager.Modify(packageDir); nil != err {
		return
	}
	// 打包
	if err = packager.Build(packageDir, outputFileName); nil != err {
		return
	}
	// 上传打包好的文件
	if err = p.DestFile.Upload(outputFileName); err != nil {
		return
	}

	// 删除打包好的文件
	if removeErr := os.Remove(outputFileName); nil != removeErr {
		log.WithFields(log.Fields{
			"fileName": outputFileName,
			"error":    removeErr,
		}).Warn("删除打包好的文件出错")
	}

	return
}

func (p *Package) replace(packageDir string) (err error) {
	for _, r := range p.Replaces {
		if err = r.Replace(packageDir); nil != err {
			break
		}
	}

	return
}
