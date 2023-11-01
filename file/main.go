package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/debeando/go-common/cast"
)

func Exist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
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

func ReadLineByLine(path string, fn func(string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fn(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
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

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func GetInt64(path string) int64 {
	lines := ReadAsString(path)
	if len(lines) > 0 {
		return cast.StringToInt64(lines)
	}
	return 0
}

func GetInt(path string) int {
	lines := ReadAsString(path)
	if len(lines) > 0 {
		return cast.StringToInt(lines)
	}
	return 0
}

func Create(f string) bool {
	if !Exist(f) {
		var file, err = os.Create(f)
		if err != nil {
			return false
		}
		defer file.Close()
	}

	return true
}

func Write(f string, s string) bool {
	// open file using READ & WRITE permission
	file, err := os.OpenFile(f, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return false
	}

	// write some text line-by-line to file
	_, err = file.WriteString(s)
	if err != nil {
		return false
	}

	// save changes
	err = file.Sync()
	if err != nil {
		return false
	}

	return true
}

func Truncate(f string) bool {
	// open file using READ & WRITE permission
	file, err := os.OpenFile(f, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return false
	}

	file.Truncate(0)
	file.Seek(0, 0)
	file.Sync()

	return true
}

func Delete(f string) bool {
	if err := os.Remove(f); err != nil {
		return false
	}
	return true
}
