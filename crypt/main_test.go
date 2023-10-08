package crypt_test

import (
	"testing"

	"github.com/debeando/go-common/crypt"
)

func TestMD5(t *testing.T) {
	expected := "098f6bcd4621d373cade4e832627b4f6"
	result := crypt.MD5("test")

	if result != expected {
		t.Error("Expected: 098f6bcd4621d373cade4e832627b4f6")
	}
}
