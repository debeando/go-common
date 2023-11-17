package main

import (
	"fmt"

	"github.com/debeando/go-common/mysql"
)

func main() {
	m := mysql.New("example", "root:admin@tcp(127.0.0.1:3306)/")
	m.Connect()
	one := m.FetchOne("SELECT NOW()")
	fmt.Printf("%s\n", one)
}
