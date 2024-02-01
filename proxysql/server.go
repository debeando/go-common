package proxysql

import (
	"fmt"

	"github.com/debeando/go-common/mysql"
)

type Server struct {
	Connection        *mysql.Connection `yaml:"-"`
	HostgroupID       uint8             `yaml:"hostgroup_id"`
	Hostname          string            `yaml:"hostname"`
	MaxConnections    uint16            `yaml:"max_connections"`
	MaxReplicationLag uint16            `yaml:"max_replication_lag"`
	Port              uint16            `yaml:"port"`
	Status            string            `yaml:"status"`
	Weight            uint16            `yaml:"weight"`
	Stats             Stats
}

func (s *Server) Insert() error {
	_, err := s.Connection.Instance.Query(s.QueryInsert())
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Update() error {
	_, err := s.Connection.Instance.Query(s.QueryUpdate())
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Fetcher() error {
	err := s.Connection.Instance.QueryRow(s.QuerySelect()).Scan(
		&s.HostgroupID,
		&s.Hostname,
		&s.Port,
		&s.Status,
		&s.Weight,
		&s.MaxConnections,
		&s.MaxReplicationLag)

	// c.ProxySQL.Servers.First().Stats.Connection = c.ProxySQL.Connection
	// c.ProxySQL.Servers.First().Stats.HostgroupID = c.ProxySQL.Servers.First().HostgroupID
	// c.ProxySQL.Servers.First().Stats.Hostname = c.AWS.RDS.Replica.Instance.Endpoint
	// c.ProxySQL.Servers.First().Stats.Fetcher()

	return err
}

func (s *Server) Delete() error {
	_, err := s.Connection.Instance.Query(s.QueryDelete())
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) IsOnLine() bool {
	return s.Status == ONLINE && s.Stats.Status == ONLINE
}

func (s *Server) QueryInsert() string {
	return fmt.Sprintf(
		"INSERT INTO mysql_servers (hostgroup_id, hostname, port, status, weight, max_connections, max_replication_lag) VALUES (%d, '%s', %d, '%s', %d, %d, %d);",
		s.HostgroupID,
		s.Hostname,
		s.Port,
		s.Status,
		s.Weight,
		s.MaxConnections,
		s.MaxReplicationLag,
	)
}

func (s *Server) QuerySelect() string {
	return fmt.Sprintf(
		"SELECT hostgroup_id, hostname , port, status, weight, max_connections, max_replication_lag "+
			"FROM mysql_servers WHERE hostgroup_id = %d AND hostname = '%s' LIMIT 1;",
		s.HostgroupID,
		s.Hostname,
	)
}

func (s *Server) QueryUpdate() string {
	return fmt.Sprintf(
		"UPDATE mysql_servers "+
			"SET hostgroup_id = %d, hostname = '%s', port = %d, status = '%s', weight = %d, max_connections = %d, max_replication_lag = %d "+
			"WHERE hostgroup_id = %d AND hostname = '%s';",
		s.HostgroupID,
		s.Hostname,
		s.Port,
		s.Status,
		s.Weight,
		s.MaxConnections,
		s.MaxReplicationLag,
		s.HostgroupID,
		s.Hostname,
	)
}

func (s *Server) QueryDelete() string {
	return fmt.Sprintf(
		"DELETE FROM mysql_servers WHERE hostgroup_id = %d AND hostname = '%s';",
		s.HostgroupID,
		s.Hostname,
	)
}
