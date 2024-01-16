package mysql_test

import (
	"os"
	"testing"

	"github.com/debeando/go-common/mysql"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	m := mysql.MySQL{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: os.Getenv("MYSQL_TEST_USER"),
		Password: os.Getenv("MYSQL_TEST_PASS"),
	}

	c := mysql.New("test", m.DSN())
	assert.Empty(t, c.Instance, nil)
	e := c.Connect()
	assert.NoError(t, e)
	assert.NotEmpty(t, c.Instance, nil)
}

func TestGet(t *testing.T) {
	c := mysql.Get("test")
	assert.NotEmpty(t, c.Instance, nil)
}

func TestClose(t *testing.T) {
	c := mysql.Get("test")
	assert.NotEmpty(t, c.Instance, nil)
	c.Close()
	assert.Empty(t, c.Instance, nil)
}
