package gaihosts

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetHostsFileByte 获取文件中的内容
func GetHostsFileByte() (res []byte,err error) {
	_, filePath := GetHostPath()
	stat, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf(ErrNoHostsFile,err)
	}
	if stat.IsDir(){
		return nil, fmt.Errorf(ErrHostsIsFolder,err)
	}

	res, err = ioutil.ReadFile(filePath)

	return res, err
}

// GetHostsByte 获取文件中的内容
func SetHostsFileByte(body []byte) (err error) {
	_, filePath := GetHostPath()
	stat, err := os.Stat(filePath)
	if err != nil {
		return  fmt.Errorf(ErrNoHostsFile,err)
	}
	if stat.IsDir(){
		return  fmt.Errorf(ErrHostsIsFolder,err)
	}

	err = ioutil.WriteFile(filePath,body,os.ModePerm)

	return  err
}
