package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"time"

	"github.com/debeando/go-common/log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type Column struct {
	Name     string
	Kind     reflect.Type
	Nullable bool
}

type Connection struct {
	Instance driver.Conn
	Name     string
}

var instance = make(map[string]*Connection)

func Instance(name string) *Connection {
	if instance[name] == nil {
		instance[name] = &Connection{}
		instance[name].Name = name
	}
	return instance[name]
}

func (c *Connection) Connect(host, port, database, username, password string) error {
	var err error

	if c.Instance != nil {
		return errors.New("ClickHouse can't connect because instance is empty.")
	}

	log.DebugWithFields("ClickHouse connection settings.", log.Fields{
		"Host":     host,
		"Port":     port,
		"Database": database,
		"Username": username,
		"Password": password,
	})

	c.Instance, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", host, port)},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		return err
	}

	if err = c.Instance.Ping(context.Background()); err != nil {
		return err
	}

	log.Debug("ClickHouse connected!")

	return nil
}

func (c *Connection) Query(query string) (map[int]map[string]any, error) {
	log.DebugWithFields("ClickHouse execute.", log.Fields{
		"Query": query,
	})

	rows, err := c.Instance.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	row_id := 0
	columns := []Column{}
	types := rows.ColumnTypes()
	dataset := make(map[int]map[string]any)

	for _, c := range types {
		columns = append(columns, Column{
			Name:     c.Name(),
			Kind:     c.ScanType(),
			Nullable: c.Nullable(),
		})
	}

	values := make([]any, len(types))

	for i := range types {
		values[i] = reflect.New(types[i].ScanType()).Interface()
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		row := make(map[string]any)

		for i, v := range values {
			switch v := v.(type) {
			case *bool:
				row[columns[i].Name] = *v
			case *string:
				row[columns[i].Name] = *v
			case *int8:
				row[columns[i].Name] = *v
			case *int16:
				row[columns[i].Name] = *v
			case *int32:
				row[columns[i].Name] = *v
			case *int64:
				row[columns[i].Name] = *v
			case *uint8:
				row[columns[i].Name] = *v
			case **uint8:
				row[columns[i].Name] = *v
			case *uint16:
				row[columns[i].Name] = *v
			case *uint32:
				row[columns[i].Name] = *v
			case *uint64:
				row[columns[i].Name] = *v
			case **big.Int:
				row[columns[i].Name] = *v
			case *float32, *float64:
				if value := *(v.(*float64)); !math.IsNaN(value) {
					row[columns[i].Name] = value
				}
			case *time.Time:
				row[columns[i].Name] = v.Format("2006-01-02 15:04:05")
			case *uuid.UUID:
				row[columns[i].Name] = *v
			default:
				row[columns[i].Name] = nil
			}
		}

		dataset[row_id] = row

		row_id++
	}

	return dataset, nil
}
