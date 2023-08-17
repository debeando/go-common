package proxysql

import (
	"fmt"

	"github.com/debeando/go-common/cast"
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
	Servers    []Server
}

type Server struct {
	HostgroupID       uint8  `json:"hostgroup_id"`
	Hostname          string `json:"hostname"`
	MaxConnections    uint16 `json:"max_connections"`
	MaxReplicationLag uint16 `json:"max_replication_lag"`
	Status            string `json:"status"`
	Weight            uint16 `json:"weight"`
}

func (p *ProxySQL) AddServer(s Server) int {
	p.Servers = append(p.Servers, s)

	return 0
}

func (p *ProxySQL) GetStatusServer(index int) string {
	sql := fmt.Sprintf(
		"SELECT status FROM mysql_servers WHERE hostgroup_id = %d AND hostname = '%s' LIMIT 1;",
		p.Servers[index].HostgroupID,
		p.Servers[index].Hostname,
	)
	result, _ := p.Connection.Query(sql)

	if len(result) == 1 {
		return result[0]["status"]
	}
	return ServerStatus
}

func (p *ProxySQL) SetStatusServer(index int, ss string) {
	sql := fmt.Sprintf(
		"UPDATE mysql_servers SET status = '%s' WHERE hostgroup_id = %d AND hostname = '%s';",
		ss,
		p.Servers[index].HostgroupID,
		p.Servers[index].Hostname,
	)
	p.Connection.Query(sql)
}

func (p *ProxySQL) ExistServer(index int) bool {
	sql := fmt.Sprintf(
		"SELECT count() AS cnt FROM mysql_servers WHERE hostgroup_id = %d AND hostname = '%s';",
		p.Servers[index].HostgroupID,
		p.Servers[index].Hostname,
	)
	result, _ := p.Connection.Query(sql)

	return ((len(result) == 1) && (cast.StringToInt(result[0]["cnt"]) == 1))
}

func (p *ProxySQL) DeleteServer(index int) {
	sql := fmt.Sprintf(
		"DELETE FROM mysql_servers WHERE hostgroup_id = %d AND hostname = '%s';",
		p.Servers[index].HostgroupID,
		p.Servers[index].Hostname,
	)
	p.Connection.Query(sql)
}

func (p *ProxySQL) InsertServer(index int) {
	sql := fmt.Sprintf(
		"INSERT INTO mysql_servers (hostgroup_id, hostname, status, max_connections, weight, max_replication_lag, comment) VALUES (%d, '%s', '%s', %d, %d, %d, 'Managed by DeBeAndo');",
		p.Servers[index].HostgroupID,
		p.Servers[index].Hostname,
		ONLINE,
		p.Servers[index].MaxConnections,
		p.Servers[index].Weight,
		p.Servers[index].MaxReplicationLag,
	)

	p.Connection.Query(sql)
}

func (p *ProxySQL) EnableServer(index int) {
	if p.ExistServer(index) && p.GetStatusServer(index) == OFFLINE_HARD {
		p.DeleteServer(index)
		p.InsertServer(index)
	}

	if p.ExistServer(index) && p.GetStatusServer(index) == OFFLINE_SOFT {
		p.SetStatusServer(index, ONLINE)
	}

	if !p.ExistServer(index) {

		p.InsertServer(index)
	}

	p.LoadServers()
	p.SaveServers()
}

func (p *ProxySQL) LoadServers() {
	p.Connection.Query("LOAD MYSQL SERVERS TO RUNTIME;")
}

func (p *ProxySQL) SaveServers() {
	p.Connection.Query("SAVE MYSQL SERVERS TO DISK;")
}

func (p *ProxySQL) StatConnectionPoolReset() {
	p.Connection.Query("SELECT * FROM stats_mysql_connection_pool_reset;")
}

func (p *ProxySQL) StatConnectionPool(index int) {
	sql := fmt.Sprintf(
		"SELECT hostgroup, substr(srv_host, 0, instr(srv_host, '.')) AS host, status, ConnUsed, ConnOK, ConnERR, Queries FROM stats_mysql_connection_pool WHERE srv_host = '%s';",
		p.Servers[index].Hostname,
	)
	p.Connection.Query(sql)
}
