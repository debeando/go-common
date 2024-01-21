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
	s.ProxySQL = p
	p.Servers.Add(s)
}
