package gaihosts

import (
	"bytes"
	"testing"
)

func TestGetHostsFileByte(t *testing.T) {

	gotRes, err := GetHostsFileByte()
	if err != nil {
		t.Errorf("GetHostsFileByte() error = %v", err)
		return
	}
	//t.Logf("%s", gotRes)

	t.Log("get success !!! ")

	if IsRoot(){
		s:=bytes.NewBufferString("\n 127.0.0.1 www.baidu.com").Bytes()
		gotRes=append(gotRes,s...)
		err=SetHostsFileByte(gotRes)
		if err != nil {
			t.Errorf("SetHostsFileByte() error = %v", err)
			return
		}
	}
}
