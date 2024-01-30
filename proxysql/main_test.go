package proxysql_test

import (
	"testing"

	"github.com/debeando/go-common/mysql"
	"github.com/debeando/go-common/proxysql"

	"github.com/stretchr/testify/assert"

	"gopkg.in/yaml.v3"
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
	assert.NotEmpty(t, p.Servers.First().Connection)

	p.Servers.Reset()
}

func TestUnmarshalYaml(t *testing.T) {
	pT := proxysql.ProxySQL{}
	cT := `
---
mysql:
  host: 127.0.0.1
  port: 6032
  username: radmin
  password: radmin
servers:
  - hostgroup_id: 20
    hostname: "127.0.0.1"
    port: 3306
    status: ONLINE
    weight: 1
    max_connections: 0
    max_replication_lag: 0
`

	err := yaml.Unmarshal([]byte(cT), &pT)

	assert.NoError(t, err)
	assert.Equal(t, pT.Servers.Count(), 1)
	assert.Equal(t, pT.Servers.First().HostgroupID, uint8(20))
	assert.Equal(t, pT.Servers.First().Hostname, "127.0.0.1")
	assert.Equal(t, pT.Servers.First().Port, uint16(3306))
	assert.Equal(t, pT.Servers.First().Status, proxysql.ONLINE)
}

func TestLink(t *testing.T) {
	z := proxysql.ProxySQL{
		MySQL: mysql.MySQL{
			Host:     "127.0.0.1",
			Port:     6032,
			Username: "radmin",
			Password: "radmin",
		},
		Servers: []proxysql.Server{
			{HostgroupID: 10},
			{HostgroupID: 11},
		},
	}

	z.Connection = mysql.New("proxysql", z.MySQL.DSN())
	z.Link()

	assert.Equal(t, z.Servers.Count(), 2)
	assert.NotEmpty(t, z.Connection, nil)
	assert.NotEmpty(t, z.Servers.First(), nil)
}
