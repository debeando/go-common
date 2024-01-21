package proxysql

type Stats struct {
	ProxySQL       *ProxySQL
	ConnectionPool ConnectionPool
}

const QueryConnectionPool = "SELECT hostgroup, srv_host, srv_port, status, ConnUsed, ConnFree, ConnOK, ConnERR, MaxConnUsed, Queries, Queries_GTID_sync, Bytes_data_sent, Bytes_data_recv, Latency_us FROM stats_mysql_connection_pool;"
const QueryConnectionPoolReset = "SELECT * FROM stats_mysql_connection_pool_reset;"

type ConnectionPool struct {
	ProxySQL        *ProxySQL `db:"-"`
	HostgroupID     uint8     `db:"hostgroup"`
	Hostname        string    `db:"srv_host"`
	Port            uint16    `db:"srv_port"`
	Status          string    `db:"status"`
	ConnUsed        uint64    `db:"ConnUsed"`
	ConnFree        uint64    `db:"ConnFree"`
	ConnOK          uint64    `db:"ConnOK"`
	ConnERR         uint64    `db:"ConnERR"`
	MaxConnUsed     uint64    `db:"MaxConnUsed"`
	Queries         uint64    `db:"Queries"`
	QueriesGTIDSync uint64    `db:"Queries_GTID_sync"`
	BytesDataSent   uint64    `db:"Bytes_data_sent"`
	BytesDataRecv   uint64    `db:"Bytes_data_recv"`
	Latency         uint64    `db:"Latency_us"`
}

func (p *ConnectionPool) Fetcher() error {
	return p.ProxySQL.Connection.Instance.QueryRow(QueryConnectionPool).Scan(
		&p.HostgroupID,
		&p.Hostname,
		&p.Port,
		&p.Status,
		&p.ConnUsed,
		&p.ConnFree,
		&p.ConnOK,
		&p.ConnERR,
		&p.MaxConnUsed,
		&p.Queries,
		&p.QueriesGTIDSync,
		&p.BytesDataSent,
		&p.BytesDataRecv,
		&p.Latency)
}

func (p *ConnectionPool) Reset() {
	p.ProxySQL.Connection.Query(QueryConnectionPoolReset)
}
