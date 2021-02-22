package nuwa

type Packager interface {
	// Init 初始化阶段，主要做一些打包前的初始化工作，比如各种变量的初始化
	Init() (err error)

	// Decode 解码阶段，主要是把源文件解包（如果是Zip文件就解压，如果是APK就反编译，以此类推）
	Decode(inputFilename string, packageDir string) (err error)

	// Modify 修改阶段，比如替换图标、包名等
	Modify(packageDir string) (err error)

	// Build 打包阶段，在修改阶段的基础上，把处理好的程序打包成最终的包（如Exe、APK、DMG以及IPA等）
	Build(packageDir string, outputFilename string) (err error)

	// Cleanup 清理阶段，做最后的清理工作，比如删除打包过程中的中间文件
	Cleanup() (err error)

	// AddCleanupPaths 添加清理路径
	AddCleanupPaths(paths ...string)
}
