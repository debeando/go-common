package proxysql_test

import (
	"testing"

	"github.com/debeando/go-common/proxysql"

	"github.com/stretchr/testify/assert"
)

func TestStatsConnectionPoolFetcher(t *testing.T) {
	p.Servers.Reset()
	p.AddServer(proxysql.Server{
		HostgroupID:       uint8(10),
		Hostname:          "127.0.0.1",
		MaxConnections:    uint16(100),
		MaxReplicationLag: uint16(60),
		Port:              uint16(3307),
		Status:            proxysql.ONLINE,
		Weight:            uint16(1),
	})
	p.Servers.First().Insert()
	p.ServersLoadToRunTime()
	p.ServersSaveToDisk()

	p.Stats.Connection = p.Connection
	p.Stats.ConnectionPool.Connection = p.Connection
	p.Stats.ConnectionPool.Fetcher()

	assert.Equal(t, p.Stats.ConnectionPool.HostgroupID, uint8(10))
	assert.Equal(t, p.Stats.ConnectionPool.Hostname, "127.0.0.1")
	assert.Equal(t, p.Stats.ConnectionPool.Port, uint16(3307))

	p.Servers.First().Delete()
	p.ServersLoadToRunTime()
	p.ServersSaveToDisk()
}
