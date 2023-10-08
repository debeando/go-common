package array_test

import (
	"testing"

	"github.com/debeando/go-common/array"
)

func TestStringIn(t *testing.T) {
	list := []string{"foo", "bar"}
	result := array.StringIn("bar", list)

	if !result {
		t.Error("Expected: false")
	}

	result = array.StringIn("test", list)

	if result {
		t.Error("Expected: true")
	}

	result = array.StringIn("", list)

	if result {
		t.Error("Expected: false")
	}
}
