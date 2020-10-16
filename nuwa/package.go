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

	// PackageReq 打包请求
	PackageReq struct {
		// PackageType 打包类型
		PackageType PackageType `json:"packageType" validate:"required,oneof=windows mac android"`
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

func (pr *PackageReq) UnmarshalJSON(data []byte) (err error) {
	type cloneType PackageReq

	rawMsg := json.RawMessage{}
	pr.Package = &rawMsg

	if err = json.Unmarshal(data, (*cloneType)(pr)); err != nil {
		return
	}

	switch pr.PackageType {
	case PackageTypeWindows:
		p := Windows{}
		if err = json.Unmarshal(rawMsg, &p); err != nil {
			return
		}
		pr.Package = p
	case PackageTypeAndroid:
		p := Android{}
		if err = json.Unmarshal(rawMsg, &p); err != nil {
			return
		}
		pr.Package = p
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

func (pr *PackageReq) srcFileName(rootPath string) (srcFileName string) {
	var name string

	switch pr.PackageType {
	case PackageTypeWindows:
		name = pr.Package.(Windows).ProductName
	case PackageTypeAndroid:
		name = pr.Package.(Android).Name[DefaultAppNameKey]
	}

	srcFileName = filepath.Join(rootPath, gox.GetFileNameWithExt(
		fmt.Sprintf("i-%s-%s", name, strconv.FormatInt(time.Now().UnixNano(), 10)),
		pr.PackageType.srcFileExt(),
	))

	return
}

func (pr *PackageReq) destFileName(rootPath string) (destFileName string) {
	var name string

	switch pr.PackageType {
	case PackageTypeWindows:
		name = pr.Package.(Windows).ProductName
	case PackageTypeAndroid:
		name = pr.Package.(Android).Name[DefaultAppNameKey]
	}

	destFileName = filepath.Join(rootPath, gox.GetFileNameWithExt(
		fmt.Sprintf("o-%s-%s", name, strconv.FormatInt(time.Now().UnixNano(), 10)),
		pr.PackageType.destFileExt(),
	))

	return
}

func (pr *PackageReq) packageDir(srcFileName string) (packageDir string) {
	packageDir = gox.GetFileDir(srcFileName)

	switch pr.PackageType {
	case PackageTypeWindows:
		packageDir = fmt.Sprintf("%s-windows", packageDir)
	case PackageTypeAndroid:
		packageDir = fmt.Sprintf("%s-android", packageDir)
	}

	return
}

func (pr *PackageReq) Build(rootPath string, packager Packager) (err error) {
	// 验证基本参数
	if err = class100.Validate.Struct(pr); nil != err {
		err = gox.NewCodeError(class100.ErrorCodeValidate, "数据验证错误", err.(validator.ValidationErrors))

		return
	}

	srcFileName := pr.srcFileName(rootPath)
	// 删除源文件，以免影响下一次打包
	defer func() {
		err = os.RemoveAll(srcFileName)
	}()
	outputFileName := pr.destFileName(rootPath)
	packageDir := pr.packageDir(srcFileName)
	// 删除打包目录，以免影响下一次打包
	defer func() {
		err = os.RemoveAll(packageDir)
	}()

	// 下载源文件
	if err = pr.SrcFile.Download(srcFileName, false); err != nil {
		return
	}
	// 准备
	if err = packager.Decode(srcFileName, packageDir); nil != err {
		return
	}
	// 处理文件替换逻辑
	if err = pr.replace(packageDir); nil != err {
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
	if err = pr.DestFile.Upload(outputFileName); err != nil {
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

func (pr *PackageReq) replace(packageDir string) (err error) {
	for _, r := range pr.Replaces {
		if err = r.Replace(packageDir); nil != err {
			break
		}
	}

	return
}
