package file_test

import (
	"os"
	"testing"

	"github.com/debeando/go-common/file"
)

var wd string

func TestMain(m *testing.M) {
	wd, _ = os.Getwd()
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
