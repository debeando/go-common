package proxysql

import (
	"github.com/debeando/go-common/mysql"
)

const (
	ServerStatus string = ""
	ONLINE              = "ONLINE"
	OFFLINE_SOFT        = "OFFLINE_SOFT"
	OFFLINE_HARD        = "OFFLINE_HARD"
	SHUNNED             = "SHUNNED"
)

type ProxySQL struct {
	mysql.MySQL
	Connection *mysql.Connection
	Servers    Servers
	Stats      Stats
}

func (p *ProxySQL) AddServer(s Server) {
	s.Connection = p.Connection
	p.Servers.Add(s)
}

func (p *ProxySQL) Link() {
	for i, _ := range p.Servers {
		p.Servers[i].Connection = p.Connection
	}

	p.Stats.Connection = p.Connection
	p.Stats.ConnectionPool.Connection = p.Connection
}
