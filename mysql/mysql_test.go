package mysql_test

import (
	"testing"

	"github.com/debeando/go-common/mysql"

	"github.com/stretchr/testify/assert"
)

func TestDSN(t *testing.T) {
	m := mysql.MySQL{}
	m.Host = "127.0.0.1"
	m.Password = "test"
	m.Port = 3306
	m.Timeout = 30
	m.Username = "sakila"

	assert.Equal(t, m.DSN(), "sakila:test@tcp(127.0.0.1:3306)/?timeout=30s")
}
