package proxysql_test

import (
	"testing"

	"github.com/debeando/go-common/proxysql"

	"github.com/stretchr/testify/assert"
)

func TestStatsConnectionPoolFetcher(t *testing.T) {
	p.Servers.Reset()
	p.AddServer(proxysql.Server{
		HostgroupID:       uint8(11),
		Hostname:          "127.0.0.1",
		MaxConnections:    uint16(100),
		MaxReplicationLag: uint16(60),
		Port:              uint16(3307),
		Status:            proxysql.OFFLINE_SOFT,
		Weight:            uint16(1),
	})
	p.Servers.First().Insert()
	p.ServersLoadToRunTime()
	p.ServersSaveToDisk()

	p.Servers.First().Stats.Connection = p.Connection
	p.Servers.First().Stats.HostgroupID = p.Servers.First().HostgroupID
	p.Servers.First().Stats.Hostname = p.Servers.First().Hostname
	p.Servers.First().Stats.Fetcher()

	assert.Equal(t, p.Servers.First().Stats.HostgroupID, uint8(11))
	assert.Equal(t, p.Servers.First().Stats.Hostname, "127.0.0.1")
	assert.Equal(t, p.Servers.First().Stats.Port, uint16(3307))
	assert.Equal(t, p.Servers.First().Stats.Status, proxysql.OFFLINE_SOFT)

	p.Servers.First().Delete()
	p.ServersLoadToRunTime()
	p.ServersSaveToDisk()
}
