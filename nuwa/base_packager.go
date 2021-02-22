package nuwa

import (
	`os`

	log `github.com/sirupsen/logrus`
)

// BasePackager 基础打包
type BasePackager struct {
	cleanupPaths []string
}

func (bp *BasePackager) AddCleanupPaths(paths ...string) {
	bp.cleanupPaths = append(bp.cleanupPaths, paths...)
}

func (bp *BasePackager) Cleanup() (err error) {
	for _, path := range bp.cleanupPaths {
		if err = os.RemoveAll(path); nil != err {
			log.WithFields(log.Fields{"filepath": path, "error": err}).Error("删除文件出错")
		}
	}

	return
}
