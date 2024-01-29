package mysql

import (
	"fmt"
)

type MySQL struct {
	Host     string `json:"host" yaml:"host"`
	Port     uint16 `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Schema   string `json:"schema" yaml:"schema"`
	Status   string `json:"status" yaml:"status"`
	Timeout  uint8  `json:"timeout" yaml:"timeout"`
}

func (m *MySQL) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/?timeout=%ds",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Timeout,
	)
}

func (m *MySQL) DSNSecret() string {
	return fmt.Sprintf(
		"%s:***@tcp(%s:%d)/?timeout=%ds",
		m.Username,
		m.Host,
		m.Port,
		m.Timeout,
	)
}
