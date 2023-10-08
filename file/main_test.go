package file_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/debeando/go-common/file"
)

var wd string

func TestMain(m *testing.M) {
	wd, _ = os.Getwd()
}

func TestGetInt64FromFile(t *testing.T) {
	expected := int64(1234567890)
	result := common.GetInt64FromFile(wd + "/../assets/tests/int64.txt")

	if result != expected {
		t.Error("Expected: int64(1234567890)")
	}

	expected = int64(0)
	result = common.GetInt64FromFile(wd + "/../assets/tests/int64.log")

	if result != expected {
		t.Error("Expected: int64(0)")
	}
}
