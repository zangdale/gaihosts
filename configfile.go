package gaihosts

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 配置 hosts 的文件信息
const (
	ConfigDir = "gaihosts"
	ConfigExt = ".buguai"
)

var (
	hostFilePath  string
	hostFileDir   string
	configHostDir string
)

func init() {
	hostFileDir, hostFilePath = GetHostPath()
	configHostDir = filepath.Join(hostFileDir, ConfigDir)
}

func GetHostFilePath() string {
	return hostFilePath
}

func GetHostFileDir() string {
	return hostFileDir
}

func GetConfigHostDir() string {
	_, err := os.Stat(configHostDir)
	if err != nil {
		err = os.Mkdir(configHostDir, 0644)
		if err != nil {
			panic(err)
		}
	}
	return configHostDir
}

// SaveHostFile 保存配置 文件
func SaveHostFile(fileName string, body []byte) error {
	err := ioutil.WriteFile(fileName, body, 0644)
	return err
}

// GetConfigFilesName 获取配置信息文件夹中的配置文件
func GetConfigFilesName() ([]string, map[string]string, error) {
	fmt.Println(configHostDir)
	var fileName []string
	configFiles := map[string]string{}
	err := filepath.Walk(GetConfigHostDir(),
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == ConfigExt {
				fileName = append(fileName, info.Name())
				configFiles[info.Name()] = path
			}
			return nil
		})
	return fileName, configFiles, err
}

// GetFileBody 获取文件的内容
func GetFileBody(fileName string) ([]byte, string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, "", err
	}
	return bytes, string(bytes), nil
}
