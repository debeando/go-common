package file_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/debeando/go-common/file"
)

var (
	wd  string
	twd string
	tf  string
)

func init() {
	ex, _ := os.Executable()
	twd = filepath.Dir(ex)
	tf = twd + "/zenit.txt"
}

func TestMain(m *testing.M) {
	wd, _ = os.Getwd()
}

func TestExist(t *testing.T) {
	if file.Exist(tf) {
		t.Error("The file exist, should be not.")
	}
}

func TestCreate(t *testing.T) {
	if !file.Create(tf) {
		t.Error("Problem to create file.")
	}

	if _, err := os.Stat(tf); os.IsNotExist(err) {
		t.Error("File not exist in: zenit.txt")
	}
}

func TestWrite(t *testing.T) {
	if !file.Write(tf, "Test 1\nTest 2") {
		t.Error("Problem to write in file.")
	}
}

func TestRead(t *testing.T) {
	result := file.ReadAsString(tf)
	expected := "Test 1\nTest 2"

	if result != expected {
		t.Errorf("Expected: '%s', got: '%s'.", expected, result)
	}
}

func TestTruncate(t *testing.T) {
	if !file.Truncate(tf) {
		t.Error("Problem to truncate file.")
	}

	if len(file.Read(tf)) != 0 {
		t.Error("Is not truncated file.")
	}
}

func TestDelete(t *testing.T) {
	if !file.Delete(tf) {
		t.Error("Problem to delete file.")
	}
}

func TestGetInt64(t *testing.T) {
	expected := int64(1234567890)
	result := file.GetInt64(wd + "/../assets/tests/int64.txt")

	if result != expected {
		t.Error("Expected: int64(1234567890)")
	}

	expected = int64(0)
	result = file.GetInt64(wd + "/../assets/tests/int64.log")

	if result != expected {
		t.Error("Expected: int64(0)")
	}
}

func TestGetInt(t *testing.T) {
	expected := 1234567890
	result := file.GetInt(wd + "/../assets/tests/int.txt")

	if result != expected {
		t.Error("Expected: 1234567890")
	}
}
