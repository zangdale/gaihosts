//+build windows

package gaihosts

import (
	"os"
	"path/filepath"
)

// GetHostPath windows 下的 hosts 的路径
func GetHostPath() (fileDir, filePath string) {
	dir := os.Getenv("windir")
	fileDir = filepath.Join(dir, "system32", "drivers", "etc")
	filePath = filepath.Join(fileDir, "hosts")
	return fileDir, filePath
}

// IsRoot 是否是 管理员
func IsRoot() bool {

	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}

	return true
}
