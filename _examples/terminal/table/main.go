package main

import (
	"fmt"
	"strings"

	"github.com/debeando/go-common/table"
	"github.com/debeando/go-common/terminal"
)

func main() {
	terminal.Reset()
	terminal.Clear()
	terminal.Flush()

	// fmt.Println(terminal.Height(), terminal.Width())

	tbl := table.New("Movie", "Year", "Rate", "Votes")
	tbl.Title("Batman movies")
	tbl.Add("The Batman", 2022, 7.8, "742K")
	tbl.Add("Batman Begins", 2005, 8.2, "1.5M")
	tbl.Add("The Dark Knight", 2008, 9.0, "2.8M")
	tbl.Add("The Dark Knight Rises", 2012, 8.4, "1.8M")
	tbl.Add("Batman", 1989, 7.5, "397K")
	tbl.Add("Batman Returns", 1992, 7.1, "321K")
	tbl.Add("Batman Forever", 1995, 5.4, "263K")
	tbl.Add("Batman & Robin", 1997, 3.8, "264K")
	tbl.Add("Batman v Superman: Dawn of Justice", 2017, 6.5, "742K")
	tbl.Add("Justice League", 2017, 6.1, "741K")
	tbl.Add("Zack Snyder's Justice League", 2021, 7.9, "427K")
	// tbl.Column(0, table.Column{Alignment: table.Right})
	tbl.Column(2, table.Column{Percentage: true, Alignment: table.Right})
	// tbl.Column(2, table.Column{ZeroFill: true, Precision:3, Scale:1})
	tbl.FilterBy(1, ">= 2000").SortBy(2).Print()
	fmt.Printf("%s\nRows filtered: %d/%d\n", strings.Repeat("-", tbl.Width()), tbl.Filtered(), tbl.Count())
	fmt.Printf("Rate = Sum: %f, Min: %f, Max: %f\n", tbl.Sum(2), tbl.Min(2), tbl.Max(2))
}
