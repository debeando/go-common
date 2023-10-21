package main

import (
	"fmt"

	"github.com/debeando/go-common/mysql"
)

func main() {
	m := mysql.New("example", "root:admin@tcp(127.0.0.1:3306)/")
	m.Connect()
	m.FetchAll("SELECT DISTINCT table_schema FROM information_schema.TABLES", func(row map[string]string) {
		fmt.Println(row)
	})

	m.FetchAll("SHOW GLOBAL STATUS", func(row map[string]string) {
		if value, ok := mysql.ParseNumberValue(row["Value"]); ok {
			fmt.Println(value)
		}
	})
}
