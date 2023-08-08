package json_test

import (
	"testing"

	"github.com/debeando/go-common/json"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	s, e := json.Escape(`{"foo":"bar"}`)
	assert.Equal(t, s, `{\"foo\":\"bar\"}`)
	assert.NoError(t, e)
}
