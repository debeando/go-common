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
