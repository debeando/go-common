package table

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
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
	AddHeader(...any) Table
	AddRow(...any) Table
	ColumnAlignment(int, Alignment) Table
	Count() int
	FilterBy(int, string) Table
	Filtered() int
	Padding(uint) Table
	Print()
	Title(string) Table
}

type Value any

type table struct {
	columnAlignment map[int]Alignment
	count           int
	filtered        int
	header          Row
	headerUpper     bool
	padding         uint
	rows            Rows
	rowsFilters     map[int]string
	title           string
	width           int
	widths          map[int]int
}

func New(header ...any) Table {
	t := table{}
	t.columnAlignment = map[int]Alignment{}
	t.rowsFilters = map[int]string{}
	t.AddHeader(header...)
	t.Padding(uint(2))

	return &t
}

func (t *table) Title(title string) Table {
	t.title = title
	return t
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
	t.count++
	return t
}

func (t *table) Padding(p uint) Table {
	t.padding = p
	return t
}

func (t *table) ColumnAlignment(index int, a Alignment) Table {
	t.columnAlignment[index] = a
	return t
}

func (t *table) FilterBy(index int, condition string) Table {
	t.rowsFilters[index] = condition

	return t
}

func (t *table) printTitle() {
	c := color.New(color.FgGreen)
	c.Println(t.title)
	c.Println(strings.Repeat("=", t.width))
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
	if t.columnAlignment[index] == Right {
		return t.lenOffset(value, t.widths[index]) + value + t.printPadding()
	}
	return value + t.lenOffset(value, t.widths[index]) + t.printPadding()
}

func (t *table) Print() {
	t.filterRows()
	t.calculateWidths()
	t.calculateWidth()
	t.printTitle()
	t.printHeader()
	t.printRows()
}

func (t *table) filterRows() {
	for index, row := range t.rows {
		for rowIndex, rowValue := range row {
			if condition, ok := t.rowsFilters[rowIndex]; ok {
				if evalCondition(condition, rowValue) {
					t.rows.Remove(index)
					t.filtered++
				}
			}
		}
	}
}

func (t *table) calculateWidths() {
	t.widths = map[int]int{}

	for headerIndex, headerValue := range t.header {
		t.widths[headerIndex] = math.Max(t.widths[headerIndex], len(fmt.Sprint(headerValue)))
	}

	for _, row := range t.rows {
		for rowIndex, rowValue := range row {
			t.widths[rowIndex] = math.Max(t.widths[rowIndex], len(fmt.Sprint(rowValue)))
		}
	}
}

func (t *table) calculateWidth() {
	for _, width := range t.widths {
		t.width = t.width + width + int(t.padding)
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

func (t *table) Filtered() int {
	return t.filtered
}

func (t *table) Count() int {
	return t.count
}

func evalCondition(condition string, value Value) bool {
	fs := token.NewFileSet()
	tv, _ := types.Eval(
		fs,
		nil,
		token.NoPos,
		fmt.Sprintf("%v %s", value, condition))

	return constant.BoolVal(tv.Value)
}

// Border('|', '+', '-')
// SortBy
// ColumnMin
// ColumnMax
// ColumnToPCT
// ColumnSum
// Summary
// Format(column, type)
// Limit
