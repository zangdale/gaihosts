package gaihosts

import "testing"

func TestGetConfigFilesName(t *testing.T) {

	name, m, err := GetConfigFilesName()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name, m)
	body, s, err := GetFileBody(m[name[0]])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(body, s)
}
