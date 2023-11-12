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

const DOWN_ARROW = "\u2193"

type Alignment int

type Table interface {
	Add(...any) Table
	ColumnAlignment(int, Alignment) Table
	Count() int
	FilterBy(int, string) Table
	Filtered() int
	Header(...any) Table
	Padding(uint) Table
	Print()
	SortBy(int) Table
	Title(string) Table
	Width() int
}

type Value any

type table struct {
	columnAlignment map[int]Alignment
	count           int
	filtered        int
	filters         map[int]string
	header          Row
	headerUpper     bool
	padding         uint
	rows            Rows
	sort            []int
	sortColumn      int
	title           string
	width           int
	widths          map[int]int
}

func New(header ...any) Table {
	t := table{}
	t.columnAlignment = map[int]Alignment{}
	t.filters = map[int]string{}
	t.Header(header...)
	t.Padding(uint(2))

	return &t
}

func (t *table) Title(title string) Table {
	t.title = title
	return t
}

func (t *table) Header(vals ...any) Table {
	for _, val := range vals {
		t.header = append(t.header, val)
	}

	return t
}

func (t *table) Add(vals ...any) Table {
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
	t.filters[index] = condition

	return t
}

func (t *table) SortBy(index int) Table {
	t.sortColumn = index
	t.header[index] = fmt.Sprintf("%s%s", t.header[index], DOWN_ARROW)

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

func (t *table) sortRows() {
	t.rows.SortBy(t.sortColumn)
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
	t.resetVariables()
	t.filterRows()
	t.calculateWidths()
	t.calculateWidth()
	t.printTitle()
	t.printHeader()
	t.sortRows()
	t.printRows()
}

func (t *table) resetVariables() {
	t.filtered = 0
	t.width = 0
	t.widths = map[int]int{}
}

func (t *table) filterRows() {
	for index := len(t.rows) - 1; index >= 0; index-- {
		for columnIndex, columnValue := range t.rows[index] {
			if condition, ok := t.filters[columnIndex]; ok {
				if !evalCondition(condition, columnValue) {
					t.rows.Remove(index)
					t.filtered++
				}
			}
		}
	}
}

func (t *table) calculateWidths() {
	for headerIndex, headerValue := range t.header {
		t.widths[headerIndex] = math.Max(t.widths[headerIndex], len(fmt.Sprint(headerValue)))
	}

	for _, row := range t.rows {
		for columnIndex, columnValue := range row {
			t.widths[columnIndex] = math.Max(t.widths[columnIndex], len(fmt.Sprint(columnValue)))
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

func (t *table) Width() int {
	return t.width
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
// ColumnMin
// ColumnMax
// ColumnToPCT
// ColumnSum
// Summary
// Format(column, type)
// Limit
// Center: marginLeft, marginTop, marginBottom, marginRight
