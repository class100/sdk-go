package nuwa

type Packager interface {
	// Decode 解码阶段，主要是把源文件解包（如果是Zip文件就解压，如果是APK就反编译，以此类推）
	Decode(inputFileName string, packageDir string) (err error)
	// Modify 修改阶段，比如替换图标、包名等
	Modify(packageDir string) (err error)
	// Build 打包阶段，在Prepare阶段的基础上，把处理好的程序打包成最终的包（如Exe、APK、DMG以及IPA等）
	Build(packageDir string, outputFileName string) (err error)
}
