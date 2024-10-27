package env_test

import (
	"testing"

	"github.com/debeando/go-common/env"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Setenv("ENV_VAR", "value1")

	assert.Equal(t, env.Get("ENV_VAR", "value2"), "value1")
	assert.Equal(t, env.Get("ENV_TMP", "value1"), "value1")
}

func TestGetInt(t *testing.T) {
	t.Setenv("ENV_VAR", "10")

	assert.Equal(t, env.GetInt("ENV_VAR", 20), 10)
	assert.Equal(t, env.GetInt("ENV_TMP", 0), 0)
}

func TestGetUInt8(t *testing.T) {
	t.Setenv("ENV_VAR", "255")

	assert.Equal(t, env.GetUInt8("ENV_VAR", uint8(10)), uint8(255))
	assert.Equal(t, env.GetUInt8("ENV_TMP", uint8(0)), uint8(0))
}

func TestGetUInt16(t *testing.T) {
	t.Setenv("ENV_VAR", "3306")

	assert.Equal(t, env.GetUInt16("ENV_VAR", uint16(10)), uint16(3306))
	assert.Equal(t, env.GetUInt16("ENV_TMP", uint16(0)), uint16(0))
}

func TestGetBool(t *testing.T) {
	t.Setenv("ENV_VAR", "true")

	assert.Equal(t, env.GetBool("ENV_VAR", false), true)
	assert.Equal(t, env.GetBool("ENV_TMP", false), false)
}

func TestExist(t *testing.T) {
	t.Setenv("ENV_VAR", "value1")

	assert.True(t, env.Exist("ENV_VAR"))
	assert.False(t, env.Exist("ENV_TMP"))
}
