package cast_test

import (
	"testing"

	"github.com/debeando/go-common/cast"
	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	assert.Equal(t, cast.StringToInt("123"), 123)
}
