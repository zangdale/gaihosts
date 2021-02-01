// +build linux

package gaihosts

import (
	"os"
	"path/filepath"
)

// GetHostPath linux 下的 hosts 的路径
func GetHostPath() (fileDir,filePath string)  {
	fileDir = "/etc"
	filePath=filepath.Join(fileDir,"hosts")
	return fileDir, filePath
}

// IsRoot 是否是 管理员
func IsRoot() bool {
	return os.Getuid()==0
}
