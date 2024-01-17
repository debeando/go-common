package mysql_test

import (
	"fmt"
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

func TestFetchOne(t *testing.T) {
	var r any
	c := mysql.Get("test")

	r = c.FetchOne("SELECT 'foo' AS one;")
	assert.NotEmpty(t, r, nil)
	assert.Equal(t, fmt.Sprintf("%s", r), "foo")

	r = c.FetchOne("SELECT 'foo' AS one, 'bar' AS two;")
	assert.Empty(t, r, nil)

	r = c.FetchOne("SELECT 'foo' AS one UNION SELECT 'bar';")
	assert.Equal(t, fmt.Sprintf("%s", r), "foo")
}

func TestFetchBool(t *testing.T) {
	var r bool
	c := mysql.Get("test")
	r = c.FetchBool("SELECT true;")
	assert.True(t, r)

	r = c.FetchBool("SELECT false;")
	assert.False(t, r)

	r = c.FetchBool("SELECT 0;")
	assert.False(t, r)

	r = c.FetchBool("SELECT 1;")
	assert.True(t, r)

	r = c.FetchBool("SELECT null;")
	assert.False(t, r)
}

func TestFetchAll(t *testing.T) {
	count := 0
	rows := []map[string]string{
		{"c1": "a", "c2": "0"},
		{"c1": "b", "c2": "1"},
		{"c1": "c", "c2": "2"},
	}

	c := mysql.Get("test")
	c.FetchAll("SELECT 'a' AS c1, 0 AS c2 UNION SELECT 'b', 1 UNION SELECT 'c', 2;", func(r map[string]string) {
		assert.Equal(t, r, rows[count])
		count++
	})

	assert.Equal(t, count, len(rows))
}

func TestClose(t *testing.T) {
	c := mysql.Get("test")
	assert.NotEmpty(t, c.Instance, nil)
	c.Close()
	assert.Empty(t, c.Instance, nil)
}
