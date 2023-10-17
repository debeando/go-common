package main

import (
	"fmt"

	"github.com/debeando/go-common/mongodb"
)

func main() {
	m := mongodb.New("example", "mongodb://admin:pass@127.0.0.1:27017")
	if err := m.Connect(); err == nil {
		databases := m.Databases()
		for _, database := range databases.Databases {
			fmt.Println(">", database.Name)
			collections := m.Collections(database.Name)
			for _, collection := range collections {
				fmt.Println(" -", collection)
				colStats := m.CollectionStats(database.Name, collection)
				fmt.Printf("%+v\n", colStats)
			}
		}

		serverstatus := m.ServerStatus()
		fmt.Printf("\nServer Status: %+v\n", serverstatus)
	}
}
