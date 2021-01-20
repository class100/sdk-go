package nuwa

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/class100/core"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/storezhang/gox"
	"github.com/storezhang/replace"
	"github.com/storezhang/transfer"

	"github.com/class100/sdk-go"
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
		MaxRetry int `default:"3" json:"maxRetry" validate:"omitempty,min=1,max=10"`
		// SrcFile 源文件
		SrcFile transfer.File `json:"srcFile" validate:"required"`
		// DestFile 打包后的文件
		DestFile transfer.File `json:"destFile" validate:"required"`
		// Replaces 文件替换
		Replaces []replace.Replace `json:"replaces" validate:"omitempty,dive"`
		// Notify 通知
		Notify Notify `json:"notify" validate:"omitempty,structonly"`
		// Packager 真正的打包者
		Packager interface{} `json:"packager" validate:"required"`
	}
)

// NewPackage 创建一个打包
func NewPackage(
	packageType PackageType,
	maxRetry int,
	srcFile transfer.File, destFile transfer.File,
	notify Notify,
	packager interface{},
	replaces ...replace.Replace,
) *Package {
	return &Package{
		Type:     packageType,
		MaxRetry: maxRetry,
		SrcFile:  srcFile,
		DestFile: destFile,
		Replaces: replaces,
		Notify:   notify,
		Packager: packager,
	}
}

func (pkg *Package) Tag(environmentType core.EnvironmentType) (tag string, err error) {
	switch pkg.Type {
	case PackageTypeWindows:
		tag = class100.TagPackageWindows
	case PackageTypeMac:
		tag = class100.TagPackageMac
	case PackageTypeAndroid:
		tag = class100.TagPackageAndroid
	case PackageTypeIOS:
		tag = class100.TagPackageIOS
	default:
		err = ErrorNotSupportPackage

		return
	}
	tag = fmt.Sprintf("%s-%s", tag, environmentType)

	return
}

func (pkg *Package) UnmarshalJSON(jsonBytes []byte) (err error) {
	type cloneType Package

	rawMsg := json.RawMessage{}
	pkg.Packager = &rawMsg

	if err = json.Unmarshal(jsonBytes, (*cloneType)(pkg)); err != nil {
		return
	}

	switch pkg.Type {
	case PackageTypeWindows:
		windows := Windows{}
		if err = json.Unmarshal(rawMsg, &windows); err != nil {
			return
		}
		pkg.Packager = windows
	case PackageTypeAndroid:
		android := Android{}
		if err = json.Unmarshal(rawMsg, &android); err != nil {
			return
		}
		pkg.Packager = android
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

func (pkg *Package) srcFileName(rootPath string) (srcFileName string) {
	var name string

	switch pkg.Type {
	case PackageTypeWindows:
		name = pkg.Packager.(Windows).ProductName
	case PackageTypeAndroid:
		name = pkg.Packager.(Android).Name[DefaultAppNameKey]
	}

	srcFileName = filepath.Join(rootPath, gox.GetFilenameWithExt(
		fmt.Sprintf("i-%s-%s", name, strconv.FormatInt(time.Now().UnixNano(), 10)),
		pkg.Type.srcFileExt(),
	))

	return
}

func (pkg *Package) Name() (name string) {
	switch pkg.Type {
	case PackageTypeWindows:
		name = pkg.Packager.(Windows).ProductName
	case PackageTypeAndroid:
		name = pkg.Packager.(Android).Name[DefaultAppNameKey]
	}

	return
}

func (pkg *Package) destFileName(rootPath string) (destFileName string) {
	var name string

	switch pkg.Type {
	case PackageTypeWindows:
		name = pkg.Packager.(Windows).ProductName
	case PackageTypeAndroid:
		name = pkg.Packager.(Android).Name[DefaultAppNameKey]
	}

	destFileName = filepath.Join(rootPath, gox.GetFilenameWithExt(
		fmt.Sprintf("o-%s-%s", name, strconv.FormatInt(time.Now().UnixNano(), 10)),
		pkg.Type.destFileExt(),
	))

	return
}

func (pkg *Package) packageDir(srcFileName string) (packageDir string) {
	packageDir = gox.GetFileDir(srcFileName)

	switch pkg.Type {
	case PackageTypeWindows:
		packageDir = fmt.Sprintf("%s-windows", packageDir)
	case PackageTypeAndroid:
		packageDir = fmt.Sprintf("%s-android", packageDir)
	}

	return
}

func (pkg *Package) Build(rootPath string, packager Packager) (err error) {
	// 验证基本参数
	if err = class100.Validate.Struct(pkg); nil != err {
		err = gox.NewCodeError(class100.ErrorCodeValidate, "数据验证错误", err.(validator.ValidationErrors))

		return
	}

	srcFilename := pkg.srcFileName(rootPath)
	// 删除源文件，以免影响下一次打包
	defer func() {
		if removeErr := os.RemoveAll(srcFilename); nil != removeErr {
			err = removeErr
		}
	}()

	outputFilename := pkg.destFileName(rootPath)
	defer func() {
		if removeErr := os.RemoveAll(outputFilename); nil != removeErr {
			err = removeErr
		}
	}()

	packageDir := pkg.packageDir(srcFilename)
	// 删除打包目录，以免影响下一次打包
	defer func() {
		if removeErr := os.RemoveAll(packageDir); nil != removeErr {
			err = removeErr
		}
	}()

	// 下载源文件
	if err = pkg.SrcFile.Download(srcFilename, false); err != nil {
		log.WithFields(log.Fields{
			"file":  pkg.SrcFile,
			"error": err,
		}).Error("下载未打包源文件出错")

		return
	}
	log.WithFields(log.Fields{"file": pkg.SrcFile}).Info("下载未打包源文件成功")

	// 准备
	if err = packager.Decode(srcFilename, packageDir); nil != err {
		log.WithFields(log.Fields{
			"file":  pkg.SrcFile,
			"error": err,
		}).Error("解码未打包源文件出错")

		return
	}
	log.WithFields(log.Fields{"file": pkg.SrcFile}).Info("解码未打包源文件成功")

	// 处理文件替换逻辑
	if err = pkg.replace(packageDir); nil != err {
		log.WithFields(log.Fields{
			"replaces": pkg.Replaces,
			"error":    err,
		}).Error("替换打包文件出错")

		return
	}
	log.WithFields(log.Fields{"replaces": pkg.Replaces}).Info("替换打包文件成功")

	// 处理应用包修改逻辑
	if err = packager.Modify(packageDir); nil != err {
		log.WithFields(log.Fields{
			"package": pkg,
			"error":   err,
		}).Error("修改打包文件出错")

		return
	}
	log.WithFields(log.Fields{"package": pkg}).Info("修改打包文件成功")

	// 打包
	if err = packager.Build(packageDir, outputFilename); nil != err {
		log.WithFields(log.Fields{
			"package": pkg,
			"error":   err,
		}).Error("编译打包文件出错")

		return
	}
	log.WithFields(log.Fields{"package": pkg}).Info("编译打包文件成功")

	// 上传打包好的文件
	if err = pkg.DestFile.Upload(outputFilename); err != nil {
		log.WithFields(log.Fields{
			"file":  pkg.DestFile,
			"error": err,
		}).Error("上传已打包文件出错")

		return
	}
	log.WithFields(log.Fields{"file": pkg.DestFile}).Info("上传已打包文件成功")

	return
}

func (pkg *Package) replace(packageDir string) (err error) {
	for _, r := range pkg.Replaces {
		if err = r.Replace(packageDir); nil != err {
			break
		}
	}

	return
}

func (pkg Package) String() string {
	jsonBytes, _ := json.MarshalIndent(pkg, "", "    ")

	return string(jsonBytes)
}
