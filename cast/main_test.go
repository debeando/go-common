package cast_test

import (
	"testing"

	"github.com/debeando/go-common/cast"
	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	assert.Equal(t, cast.StringToInt("123"), 123)
}

func TestStringToInt64(t *testing.T) {
	expected := int64(1234)
	result := cast.StringToInt64("1234")

	if result != expected {
		t.Error("Expected: int64(1234)")
	}

	result = cast.StringToInt64("abc")

	if result != 0 {
		t.Error("Expected: 0")
	}

	result = cast.StringToInt64("")

	if result != 0 {
		t.Error("Expected: 0")
	}
}

func TestToDateTime(t *testing.T) {
	expected := "2018-12-31 15:04:05"
	result := cast.ToDateTime("2018-12-31T15:04:05 UTC", "2006-01-02T15:04:05 UTC")

	if result != expected {
		t.Error("Expected: 2018-12-31 15:04:05")
	}
}

func TestInterfaceIsNumber(t *testing.T) {
	assert.Equal(t, cast.InterfaceIsNumber("123"), false)
	assert.Equal(t, cast.InterfaceIsNumber("abc"), false)
	assert.Equal(t, cast.InterfaceIsNumber("1bc"), false)
	assert.Equal(t, cast.InterfaceIsNumber(int(123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(int8(-123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(int32(123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(int64(-123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(uint(123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(uint8(123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(uint32(123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(uint64(123)), true)
	assert.Equal(t, cast.InterfaceIsNumber(float32(12.3)), true)
	assert.Equal(t, cast.InterfaceIsNumber(float64(12.3)), true)
}
