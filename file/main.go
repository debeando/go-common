package file

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/debeando/go-common/cast"
)

func Exist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Read(path string) []byte {
	c, _ := ioutil.ReadFile(path)
	return c
}

func ReadAsString(path string) string {
	return string(Read(path))
}

func ReadExpandEnv(path string) []byte {
	return []byte(os.ExpandEnv(ReadAsString(path)))
}

func ReadExpandEnvAsString(path string) string {
	return string(ReadExpandEnv(path))
}

func Name(n string) string {
	return n[:len(n)-len(filepath.Ext(n))]
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func GetInt64(path string) int64 {
	lines := ReadAsString(path)
	if len(lines) > 0 {
		return cast.StringToInt64(lines)
	}
	return 0
}
