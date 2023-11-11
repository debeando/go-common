package main

import (
	"fmt"

	"github.com/debeando/go-common/table"
	"github.com/debeando/go-common/terminal"
)

func main() {
	terminal.Reset()
	terminal.Clear()
	terminal.Flush()

	// fmt.Println(strings.Repeat("=", terminal.Width()))
	// fmt.Println(terminal.Height(), terminal.Width())

	tbl := table.New("Movie", "Year", "Rate", "Votes")
	tbl.Title("Batman movies")
	tbl.AddRow("The Batman",                         2022, 7.8, "742K")
	tbl.AddRow("Batman Begins",                      2005, 8.2, "1.5M")
	tbl.AddRow("The Dark Knight",                    2008, 9.0, "2.8M")
	tbl.AddRow("The Dark Knight Rises",              2012, 8.4, "1.8M")
	tbl.AddRow("Batman",                             1989, 7.5, "397K")
	tbl.AddRow("Batman Returns",                     1992, 7.1, "321K")
	tbl.AddRow("Batman Forever",                     1995, 5.4, "263K")
	tbl.AddRow("Batman & Robin",                     1997, 3.8, "264K")
	tbl.AddRow("Batman v Superman: Dawn of Justice", 2017, 6.5, "742K")
	tbl.AddRow("Justice League",                     2017, 6.1, "741K")
	tbl.AddRow("Zack Snyder's Justice League",       2021, 7.9, "427K")

	// tbl.ColumnAlignment(0, table.Right)
	// tbl.ColumnAlignment(2, table.Right)
	tbl.FilterBy(1, ">= 2000")
	tbl.Print()
	fmt.Printf("Rows filtered: %d/%d\n", tbl.Filtered(), tbl.Count())
}
