package mysql

import (
	"database/sql"
	"errors"

	"github.com/debeando/go-common/log"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Instance *sql.DB
	Name     string
	DSN      string
}

var instance = make(map[string]*Connection)

func New(name, dsn string) *Connection {
	if instance[name] == nil {
		instance[name] = &Connection{
			Name: name,
			DSN:  dsn,
		}
		instance[name].Name = name
	}
	return instance[name]
}

func Get(name string) *Connection {
	return instance[name]
}

func (c *Connection) Connect() error {
	if c.Instance == nil {
		conn, err := sql.Open("mysql", c.DSN)
		if err != nil {
			log.ErrorWithFields("MySQL:Connect", log.Fields{"name": c.Name, "message": err})
			return err
		}

		if err := conn.Ping(); err != nil {
			log.ErrorWithFields("MySQL:Connect:Ping", log.Fields{"name": c.Name, "message": err})
			return err
		}

		c.Instance = conn
	}
	return nil
}

func (c *Connection) Query(query string) (map[int]map[string]string, error) {
	if c.Instance == nil {
		return nil, errors.New("The instance is empty.")
	}

	log.DebugWithFields("MySQL:Query", log.Fields{
		"name":  c.Name,
		"query": query,
	})

	if err := c.Instance.Ping(); err != nil {
		log.ErrorWithFields("MySQL:Query:Ping", log.Fields{"name": c.Name, "message": err})
		return nil, err
	}

	// Execute the query
	rows, err := c.Instance.Query(query)
	if err != nil {
		log.ErrorWithFields("MySQL:Query", log.Fields{"name": c.Name, "message": err})
		return nil, err
	}
	defer rows.Close()

	// Get column names
	cols, _ := rows.Columns()
	if err != nil {
		log.ErrorWithFields("MySQL:Query:Columns", log.Fields{"name": c.Name, "message": err})
		return nil, err
	}

	row_id := 0
	dataset := make(map[int]map[string]string)
	columns := make([]sql.RawBytes, len(cols))
	columnPointers := make([]interface{}, len(cols))

	for i := range cols {
		columnPointers[i] = &columns[i]
	}

	for rows.Next() {
		err = rows.Scan(columnPointers...)
		if err != nil {
			log.ErrorWithFields("MySQL:Query:Scan", log.Fields{"name": c.Name, "message": err})
			return nil, err
		}

		row := make(map[string]string)

		for columnIndex, columnValue := range columns {
			row[cols[columnIndex]] = string(columnValue)
			dataset[row_id] = row
		}

		row_id++
	}

	return dataset, nil
}

func (c *Connection) FetchBool(query string) bool {
	r := c.FetchOne(query)

	if r == nil {
		return false
	}

	return true
}

func (c *Connection) FetchOne(query string) any {
	if c.Instance == nil {
		return nil
	}

	log.DebugWithFields("MySQL:FetchOne", log.Fields{
		"name":  c.Name,
		"query": query,
	})

	if err := c.Instance.Ping(); err != nil {
		log.ErrorWithFields("MySQL:FetchOne:Ping", log.Fields{"name": c.Name, "message": err})
		return nil
	}

	var val any
	row := c.Instance.QueryRow(query)
	row.Scan(&val)

	return val
}

func (c *Connection) FetchAll(query string, fn func(map[string]string)) error {
	if c.Instance == nil {
		return errors.New("The instance is empty.")
	}

	log.DebugWithFields("MySQL:FetchAll", log.Fields{
		"name":  c.Name,
		"query": query,
	})

	if err := c.Instance.Ping(); err != nil {
		log.ErrorWithFields("MySQL:FetchAll:Ping", log.Fields{"name": c.Name, "message": err})
		return err
	}

	// Execute the query
	rows, err := c.Instance.Query(query)
	if err != nil {
		log.ErrorWithFields("MySQL:FetchAll:Query", log.Fields{"name": c.Name, "message": err})
		return err
	}
	defer rows.Close()

	// Get column names
	cols, _ := rows.Columns()
	if err != nil {
		log.ErrorWithFields("MySQL:FetchAll:Columns", log.Fields{"name": c.Name, "message": err})
		return err
	}

	columns := make([]sql.RawBytes, len(cols))
	columnPointers := make([]interface{}, len(cols))

	for i := range cols {
		columnPointers[i] = &columns[i]
	}

	for rows.Next() {
		err = rows.Scan(columnPointers...)
		if err != nil {
			log.ErrorWithFields("MySQL:FetchAll:Scan", log.Fields{"name": c.Name, "message": err})
			return err
		}

		row := make(map[string]string)

		for columnIndex, columnValue := range columns {
			row[cols[columnIndex]] = string(columnValue)
		}

		fn(row)
	}

	return nil
}

func (c *Connection) Close() {
	if c.Instance != nil {
		c.Instance.Close()
		c.Instance = nil
	}
}
