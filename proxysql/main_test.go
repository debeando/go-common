package proxysql_test

import (
	"testing"

	"github.com/debeando/go-common/mysql"
	"github.com/debeando/go-common/proxysql"

	"github.com/stretchr/testify/assert"
)

var p = proxysql.ProxySQL{}

func TestConnection(t *testing.T) {
	p = proxysql.ProxySQL{
		MySQL: mysql.MySQL{
			Host:     "127.0.0.1",
			Port:     6032,
			Username: "radmin",
			Password: "radmin",
		},
	}

	p.Connection = mysql.New("proxysql", p.MySQL.DSN())

	assert.NoError(t, p.Connection.Connect())
}

func TestAddServer(t *testing.T) {
	p.Servers.Reset()
	p.AddServer(proxysql.Server{
		HostgroupID: 0,
		Hostname:    "127.0.0.1",
	})

	assert.Equal(t, p.Servers.Count(), 1)
	assert.NotEmpty(t, p.Servers.First().ProxySQL)

	p.Servers.Reset()
}
