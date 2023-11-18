package table

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

const DOWN_ARROW = "\u2193"

const (
	Left Alignment = iota
	Center
	Right
)

type Value any
type Alignment int

type Stats struct {
	Min float64
	Max float64
	Sum float64
	Len int
}

type Table interface {
	Add(...any) Table
	Column(int, Column) Table
	Count() int
	FilterBy(int, string) Table
	Filtered() int
	Max(int) float64
	Min(int) float64
	Padding(uint) Table
	Print()
	SortBy(int) Table
	Sum(int) float64
	Title(string) Table
	Width() int
}

type table struct {
	columns    map[int]Column
	count      int
	filtered   int
	filters    map[int]string
	padding    uint
	rows       Rows
	sort       []int
	sortColumn int
	stats      map[int]Stats
	title      string
	width      int
}

func New() Table {
	t := table{}
	t.filters = map[int]string{}
	t.columns = map[int]Column{}
	t.stats = map[int]Stats{}
	t.padding = uint(2)

	return &t
}

func (t *table) Title(title string) Table {
	t.title = title
	return t
}

func (t *table) Add(vals ...any) Table {
	row := Row{}
	for _, val := range vals {
		row = append(row, Field{Value: val})
	}
	t.rows = append(t.rows, row)
	t.count++
	return t
}

func (t *table) Padding(p uint) Table {
	t.padding = p
	return t
}

func (t *table) Column(index int, column Column) Table {
	t.columns[index] = column
	return t
}

func (t *table) FilterBy(index int, condition string) Table {
	t.filters[index] = condition

	return t
}

func (t *table) SortBy(index int) Table {
	t.sortColumn = index
	t.SetColumnName(
		index,
		fmt.Sprintf("%s%s", t.columns[index].Name, DOWN_ARROW))

	return t
}

func (t *table) SetColumnName(index int, name string) {
	tmp := t.columns[index]
	tmp.SetName(name)
	(*t).columns[index] = tmp
}

func (t *table) Print() {
	t.resetVariables()
	t.clearColumnsValues()
	t.filterRows()
	t.calculateColumnStats()
	t.calculateTableWidth()
	t.printTitle()
	t.printHeader()
	t.sortRows()
	t.printRows()
}

func (t *table) Filtered() int {
	return t.filtered
}

func (t *table) Count() int {
	return t.count
}

func (t *table) Min(index int) float64 {
	return t.stats[index].Min
}

func (t *table) Max(index int) float64 {
	return t.stats[index].Max
}

func (t *table) Sum(index int) float64 {
	return t.stats[index].Sum
}

func (t *table) Width() int {
	return t.width
}

func (t *table) printTitle() {
	c := color.New(color.FgGreen)
	c.Println(t.title)
	c.Println(strings.Repeat("=", t.width))
}

func (t *table) printHeader() {
	c := color.New(color.FgGreen).Add(color.Underline)
	c.Println(t.buildHeader())
}

func (t *table) printRows() {
	for index, row := range t.rows {
		fmt.Println(t.buildRow(index, row))
	}
}

func (t *table) sortRows() {
	t.rows.SortBy(t.sortColumn)
}

func (t *table) buildHeader() (p string) {
	for i := 0; i < len(t.columns); i++ {
		if t.columns[i].Alignment == Right {
			p = p + t.lenOffset(t.columns[i].Name, t.stats[i].Len) + t.columns[i].Name + t.printPadding()
		} else {
			p = p + t.columns[i].Name + t.lenOffset(t.columns[i].Name, t.stats[i].Len) + t.printPadding()
		}
	}
	return p
}

func (t *table) buildRow(rowIndex int, row Row) (p string) {
	for columnIndex, columnValue := range row {
		p = p + t.buildColumn(rowIndex, columnIndex, columnValue)
	}
	return p
}

func (t *table) buildColumn(rowIndex, columnIndex int, columnValue Field) (field string) {
	field = columnValue.ToString()
	if rowIndex >= 0 && t.columns[columnIndex].Percentage == true {
		field = fmt.Sprintf("%d%%", columnValue.Percentage(t.Min(columnIndex), t.Max(columnIndex)))
	}

	if rowIndex >= 0 && t.columns[columnIndex].ZeroFill == true {
		field = columnValue.ZeroFill(
			t.columns[columnIndex].Precision,
			t.columns[columnIndex].Scale)
	}

	if t.columns[columnIndex].Alignment == Right {
		return t.lenOffset(field, t.stats[columnIndex].Len) + field + t.printPadding()
	}

	return field + t.lenOffset(field, t.stats[columnIndex].Len) + t.printPadding()
}

func (t *table) resetVariables() {
	t.filtered = 0
	t.width = 0
}

func (t *table) clearColumnsValues() {
	for rowIndex, rowValue := range t.rows {
		for columnIndex, _ := range rowValue {
			(*t).rows[rowIndex][columnIndex].Clear()
			(*t).rows[rowIndex][columnIndex].Truncate(t.columns[columnIndex].Truncate)
		}
	}
}

func (t *table) filterRows() {
	for index := len(t.rows) - 1; index >= 0; index-- {
		for columnIndex, columnValue := range t.rows[index] {
			if condition, ok := t.filters[columnIndex]; ok {
				if !columnValue.EvalCondition(condition) {
					t.rows.Remove(index)
					t.filtered++
				}
			}
		}
	}
}

func (t *table) calculateColumnStats() {
	for index, value := range t.columns {
		var v float64

		if len(t.rows) > 0 {
			v = t.rows[0][index].ToFloat64()
		}

		t.stats[index] = Stats{
			Len: t.calculateColumnLen(index, value.GetWidth()),
			Min: t.calculateColumnMin(index, v),
			Max: t.calculateColumnMax(index, v),
			Sum: t.calculateColumnSum(index),
		}
	}
}

func (t *table) calculateColumnMin(index int, value float64) float64 {
	for x := 1; x < len(t.rows); x++ {
		value = math.Min(value, t.rows[x][index].ToFloat64())
	}

	return value
}

func (t *table) calculateColumnMax(index int, value float64) float64 {
	for x := 1; x < len(t.rows); x++ {
		value = math.Max(value, t.rows[x][index].ToFloat64())
	}

	return value
}

func (t *table) calculateColumnSum(index int) (sum float64) {
	for x := 0; x < len(t.rows); x++ {
		sum = sum + t.rows[x][index].ToFloat64()
	}

	return sum
}

func (t *table) calculateColumnLen(index int, value int) int {
	for x := 0; x < len(t.rows); x++ {
		value = int(math.Max(float64(value), float64(t.rows[x][index].Len())))
	}

	return value
}

func (t *table) calculateTableWidth() {
	for _, stats := range t.stats {
		t.width = t.width + stats.Len + int(t.padding)
	}
}

func (t *table) lenOffset(s string, w int) string {
	l := w - utf8.RuneCountInString(s)
	if l < 0 {
		return ""
	}

	return strings.Repeat(" ", l)
}

func (t *table) printPadding() string {
	return strings.Repeat(" ", int(t.padding))
}
