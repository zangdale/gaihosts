package gaihosts

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetHostsFileByte 获取文件中的内容
func GetHostsFileByte() (res []byte,err error) {
	stat, err := os.Stat(hostFilePath)
	if err != nil {
		return nil, fmt.Errorf(ErrNoHostsFile,err)
	}
	if stat.IsDir(){
		return nil, fmt.Errorf(ErrHostsIsFolder,err)
	}

	res, err = ioutil.ReadFile(hostFilePath)

	return res, err
}

// GetHostsByte 获取文件中的内容
func SetHostsFileByte(body []byte) (err error) {

	stat, err := os.Stat(hostFilePath)
	if err != nil {
		return  fmt.Errorf(ErrNoHostsFile,err)
	}
	if stat.IsDir(){
		return  fmt.Errorf(ErrHostsIsFolder,err)
	}

	err = ioutil.WriteFile(hostFilePath,body,os.ModePerm)

	return  err
}
