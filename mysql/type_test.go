package mysql_test

import (
	"testing"

	"github.com/debeando/go-common/mysql"

	"github.com/stretchr/testify/assert"
)

func TestParseValue(t *testing.T) {
	var ok bool
	var value int64

	value, ok = mysql.ParseNumberValue("yes")
	assert.True(t, ok)
	assert.Equal(t, value, int64(1))

	value, ok = mysql.ParseNumberValue("Yes")
	assert.True(t, ok)
	assert.Equal(t, value, int64(1))

	value, ok = mysql.ParseNumberValue("YES")
	assert.True(t, ok)
	assert.Equal(t, value, int64(1))

	value, ok = mysql.ParseNumberValue("no")
	assert.True(t, ok)
	assert.Equal(t, value, int64(0))

	value, ok = mysql.ParseNumberValue("No")
	assert.True(t, ok)
	assert.Equal(t, value, int64(0))

	value, ok = mysql.ParseNumberValue("NO")
	assert.True(t, ok)
	assert.Equal(t, value, int64(0))

	value, ok = mysql.ParseNumberValue("ON")
	assert.True(t, ok)
	assert.Equal(t, value, int64(1))

	value, ok = mysql.ParseNumberValue("OFF")
	assert.True(t, ok)
	assert.Equal(t, value, int64(0))

	value, ok = mysql.ParseNumberValue("true")
	assert.False(t, ok)
	assert.Equal(t, value, int64(0))

	value, ok = mysql.ParseNumberValue("1234567890")
	assert.True(t, ok)
	assert.Equal(t, value, int64(1234567890))

}

func TestClearUser(t *testing.T) {
	user := "test[test] @ [127.0.0.1]"
	result := mysql.ClearUser(user)

	assert.Equal(t, result, "test")
}
