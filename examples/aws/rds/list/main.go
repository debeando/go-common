package main

import (
	"github.com/debeando/go-common/aws/rds"
	"github.com/rodaine/table"
)

func main() {
	r := rds.Config{}
	r.Init()
	instances := r.List()

	tbl := table.New("Engine", "Version", "Identifier", "Class", "Status")
  for _, instance := range instances {
    tbl.AddRow(instance.Engine, instance.Version, instance.Identifier, instance.Class, instance.Status)
  }
	tbl.Print()
}
