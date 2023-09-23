package file

import (
	"io/ioutil"
	"path/filepath"
)

func Read(path string) []byte {
	c, _ := ioutil.ReadFile(path)
	return c
}

func ReadAsString(path string) string {
	return string(Read(path))
}

func Name(n string) string {
	return n[:len(n)-len(filepath.Ext(n))]
}

func Dir(path string) string {
	return filepath.Dir(path)
}
