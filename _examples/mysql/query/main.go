package main

import (
	"fmt"

	"github.com/debeando/go-common/mysql"
)

func main() {
	m := mysql.New("example", "root:admin@tcp(127.0.0.1:3306)/")
	m.Connect()
	rows, _ := m.Query("SELECT DISTINCT table_schema FROM information_schema.TABLES")
	fmt.Println(rows)
}
