package proxysql_test

import (
	"testing"

	"github.com/debeando/go-common/proxysql"

	"github.com/stretchr/testify/assert"
)

func TestServersReset(t *testing.T) {
	p.Servers.Reset()

	assert.Equal(t, p.Servers.Count(), 0)
}

func TestServersAdd(t *testing.T) {
	p.Servers.Reset()
	p.Servers.Add(proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
	})

	assert.Equal(t, len(p.Servers), 1)
}

func TestServersCount(t *testing.T) {
	p.Servers.Reset()
	p.Servers.Add(proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
	})

	p.Servers.Add(proxysql.Server{
		HostgroupID: 11,
		Hostname:    "127.0.0.1",
	})

	assert.Equal(t, p.Servers.Count(), 2)
}

func TestServersFirst(t *testing.T) {
	p.Servers.Reset()
	p.Servers.Add(proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
	})

	assert.Equal(t, p.Servers.First().HostgroupID, uint8(10))
	assert.Equal(t, p.Servers.First().Hostname, "127.0.0.1")
}

func TestServersGet(t *testing.T) {
	p.Servers.Reset()
	p.Servers.Add(proxysql.Server{
		HostgroupID: 11,
		Hostname:    "127.0.0.1",
	})

	s := p.Servers.Get(uint8(11), "127.0.0.1")

	assert.Equal(t, s.HostgroupID, uint8(11))
	assert.Equal(t, s.Hostname, "127.0.0.1")

	e := p.Servers.Get(uint8(12), "127.0.0.2")

	assert.Equal(t, e.HostgroupID, uint8(0))
	assert.Equal(t, e.Hostname, "")
}
