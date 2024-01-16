package proxysql

import (
	"errors"
	"fmt"

	"github.com/debeando/go-common/cast"
	"github.com/debeando/go-common/mysql"
	"github.com/debeando/go-common/time"
)

const (
	ServerStatus string = ""
	ONLINE              = "ONLINE"
	OFFLINE_SOFT        = "OFFLINE_SOFT"
	OFFLINE_HARD        = "OFFLINE_HARD"
	SHUNNED             = "SHUNNED"
)

type ProxySQL struct {
	Connection *mysql.Connection
	Host       string   `yaml:"host"`
	Password   string   `yaml:"password"`
	Port       uint16   `yaml:"port"`
	Schema     string   `yaml:"schema"`
	Servers    []Server `yaml:"servers"`
	Status     string   `yaml:"status"`
	Timeout    uint8    `yaml:"timeout"`
	Username   string   `yaml:"username"`
}

type Server struct {
	HostgroupID       uint8  `yaml:"hostgroup_id"`
	Hostname          string `yaml:"hostname"`
	MaxConnections    uint16 `yaml:"max_connections"`
	MaxReplicationLag uint16 `yaml:"max_replication_lag"`
	Status            string `yaml:"status"`
	Weight            uint16 `yaml:"weight"`
}

func (p *ProxySQL) AddServer(s Server) int {
	p.Servers = append(p.Servers, s)

	return 0
}

func (p *ProxySQL) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/?timeout=%ds",
		p.Username,
		p.Password,
		p.Host,
		p.Port,
		p.Timeout,
	)
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

func (p *ProxySQL) DisableServer(index int) error {
	var cntQueries int
	var stats map[string]string

	p.SetStatusServer(0, OFFLINE_SOFT)
	p.LoadServers()
	p.SaveServers()

	time.Sleep(300000)

	p.StatConnectionPoolReset()
	stats = p.StatConnectionPool(0)
	cntQueries += cast.StringToInt(stats["Queries"])

	time.Sleep(60000)

	stats = p.StatConnectionPool(0)
	cntQueries += cast.StringToInt(stats["Queries"])

	if cntQueries > 0 {
		return errors.New("Active connections on MySQL replica.")
	}
	return nil
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

func (p *ProxySQL) StatConnectionPool(index int) map[string]string {
	sql := fmt.Sprintf(
		"SELECT hostgroup, substr(srv_host, 0, instr(srv_host, '.')) AS host, status, ConnUsed, ConnOK, ConnERR, Queries FROM stats_mysql_connection_pool WHERE srv_host = '%s';",
		p.Servers[index].Hostname,
	)
	result, _ := p.Connection.Query(sql)
	if len(result) == 1 {
		return result[0]
	}
	return nil
}
