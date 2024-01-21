package proxysql

type Servers []Server

func (ss *Servers) Reset() {
	*ss = (*ss)[:0]
}

func (ss *Servers) Add(s Server) {
	*ss = append(*ss, s)
}

func (ss *Servers) Count() int {
	return len(*ss)
}

func (ss *Servers) First() *Server {
	if ss.Count() > 0 {
		return &(*ss)[0]
	}
	return &Server{}
}

func (ss *Servers) Get(hostgroupID uint8, hostname string) *Server {
	for i, s := range *ss {
		if s.HostgroupID == hostgroupID && s.Hostname == hostname {
			return &(*ss)[i]
		}
	}
	return &Server{}
}

func (p *ProxySQL) ServersLoadToRunTime() {
	p.Connection.Query("LOAD MYSQL SERVERS TO RUNTIME;")
}

func (p *ProxySQL) ServersSaveToDisk() {
	p.Connection.Query("SAVE MYSQL SERVERS TO DISK;")
}
