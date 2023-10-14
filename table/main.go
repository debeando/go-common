package table

import (
	"fmt"
	"strings"

	"github.com/debeando/go-common/math"

	"github.com/fatih/color"
)

const (
	Left Alignment = iota
	Center
	Right
)

type Alignment int

type Table interface {
	AddRow(...any) Table
	AddHeader(...any) Table
	Padding(uint) Table
	SetFirstColumnAlignment(Alignment) Table
	Print()
}

type Value any
type Row []Value
type Rows []Row

type table struct {
	firstColumnAlignment Alignment
	header               Row
	headerAlignment      Alignment
	headerUpper          bool
	padding              uint
	rows                 Rows
	widths               map[int]int
}

func New(header ...any) Table {
	t := table{}
	t.AddHeader(header...)
	t.Padding(uint(2))

	return &t
}

func (t *table) AddHeader(vals ...any) Table {
	for _, val := range vals {
		t.header = append(t.header, val)
	}

	return t
}

func (t *table) AddRow(vals ...any) Table {
	row := Row{}
	for _, val := range vals {
		row = append(row, val)
	}
	t.rows = append(t.rows, row)
	return t
}

func (t *table) Padding(p uint) Table {
	t.padding = p
	return t
}

func (t *table) SetFirstColumnAlignment(a Alignment) Table {
	t.firstColumnAlignment = a
	return t
}

func (t *table) printHeader() {
	c := color.New(color.FgGreen).Add(color.Underline)
	c.Println(t.printRow(t.header))
}

func (t *table) printRows() {
	for _, row := range t.rows {
		fmt.Println(t.printRow(row))
	}
}

func (t *table) printRow(row Row) (p string) {
	for index, value := range row {
		p = p + t.buildRow(index, fmt.Sprint(value))
	}
	return p
}

func (t *table) buildRow(index int, value string) string {
	if index == 0 && t.firstColumnAlignment == Right {
		return t.lenOffset(value, t.widths[index]) + value + t.printPadding()
	}
	return value + t.lenOffset(value, t.widths[index]) + t.printPadding()
}

func (t *table) Print() {
	t.calculateWidths()
	t.printHeader()
	t.printRows()
}

func (t *table) calculateWidths() {
	t.widths = map[int]int{}

	for headerIndex, headerValue := range t.header {
		t.widths[headerIndex] = math.Max(t.widths[headerIndex], len(fmt.Sprint(headerValue)))
	}

	for _, row := range t.rows {
		for columnIndex, columnValue := range row {
			t.widths[columnIndex] = math.Max(t.widths[columnIndex], len(fmt.Sprint(columnValue)))
		}
	}
}

func (t *table) lenOffset(s string, w int) string {
	l := w - len(s)
	if l < 0 {
		return ""
	}

	return strings.Repeat(" ", l)
}

func (t *table) printPadding() string {
	return strings.Repeat(" ", int(t.padding))
}
