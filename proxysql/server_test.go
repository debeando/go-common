package proxysql_test

import (
	"testing"

	"github.com/debeando/go-common/proxysql"

	"github.com/stretchr/testify/assert"
)

func TestServerSave(t *testing.T) {
	s := proxysql.Server{
		HostgroupID:       uint8(30),
		Hostname:          "127.0.0.1",
		MaxConnections:    uint16(100),
		MaxReplicationLag: uint16(60),
		Port:              uint16(3307),
		Status:            proxysql.ONLINE,
		Weight:            uint16(1),
	}
	s.ProxySQL = &p
	assert.NoError(t, s.Save())
}

func TestServerUpdate(t *testing.T) {
	s := proxysql.Server{
		HostgroupID:       uint8(30),
		Hostname:          "127.0.0.1",
		MaxConnections:    uint16(100),
		MaxReplicationLag: uint16(60),
		Port:              uint16(3307),
		Status:            proxysql.OFFLINE_SOFT,
		Weight:            uint16(1),
	}
	s.ProxySQL = &p
	assert.NoError(t, s.Update())
}

func TestServerFetcher(t *testing.T) {
	p.AddServer(proxysql.Server{
		HostgroupID: 30,
		Hostname:    "127.0.0.1",
	})

	s := p.Servers.First()
	assert.NoError(t, s.Fetcher())
	assert.Equal(t, s.HostgroupID, uint8(30))
	assert.Equal(t, s.Hostname, "127.0.0.1")
	assert.Equal(t, s.MaxConnections, uint16(100))
	assert.Equal(t, s.MaxReplicationLag, uint16(60))
	assert.Equal(t, s.Port, uint16(3307))
	assert.Equal(t, s.Status, "OFFLINE_SOFT")
	assert.Equal(t, s.Weight, uint16(1))
}

func TestServerDelete(t *testing.T) {
	s := proxysql.Server{
		HostgroupID: uint8(30),
		Hostname:    "127.0.0.1",
	}
	s.ProxySQL = &p
	assert.NoError(t, s.Delete())
}

func TestServerQueryInsert(t *testing.T) {
	s := proxysql.Server{
		HostgroupID:       uint8(30),
		Hostname:          "127.0.0.1",
		MaxConnections:    uint16(100),
		MaxReplicationLag: uint16(60),
		Port:              uint16(3307),
		Status:            proxysql.ONLINE,
	}

	assert.Equal(t,
		s.QueryInsert(),
		"INSERT INTO mysql_servers (hostgroup_id, hostname, port, status, weight, max_connections, max_replication_lag) VALUES (30, '127.0.0.1', 3307, 'ONLINE', 0, 100, 60);",
	)
}

func TestServerQuerySelect(t *testing.T) {
	s := proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
	}

	assert.Equal(t,
		s.QuerySelect(),
		"SELECT hostgroup_id, hostname , port, status, weight, max_connections, max_replication_lag "+
			"FROM mysql_servers WHERE hostgroup_id = 10 AND hostname = '127.0.0.1' LIMIT 1;")
}

func TestServerQueryUpdate(t *testing.T) {
	s := proxysql.Server{
		HostgroupID:       uint8(30),
		Hostname:          "127.0.0.1",
		MaxConnections:    uint16(100),
		MaxReplicationLag: uint16(60),
		Port:              uint16(3307),
		Status:            proxysql.ONLINE,
	}

	assert.Equal(t,
		s.QueryUpdate(),
		"UPDATE mysql_servers SET hostgroup_id = 30, hostname = '127.0.0.1', port = 3307, status = 'ONLINE', weight = 0, max_connections = 100, max_replication_lag = 60 WHERE hostgroup_id = 30 AND hostname = '127.0.0.1';")
}

func TestServerQueryDelete(t *testing.T) {
	s := proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
	}

	assert.Equal(t,
		s.QueryDelete(),
		"DELETE FROM mysql_servers WHERE hostgroup_id = 10 AND hostname = '127.0.0.1';")
}
