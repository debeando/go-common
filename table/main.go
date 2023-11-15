package table

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/debeando/go-common/cast"

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

type Column struct {
	Alignment Alignment
	// Header    string
	// Index     int
	// Filter string
	// Sort   bool
	Percentage bool
	Precision  int
	Scale      int
	ZeroFill   bool
}

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
	Header(...any) Table
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
	header     Row
	padding    uint
	rows       Rows
	sort       []int
	sortColumn int
	stats      map[int]Stats
	title      string
	width      int
}

func New(header ...any) Table {
	t := table{}
	t.filters = map[int]string{}
	t.columns = map[int]Column{}
	t.stats = map[int]Stats{}
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
	t.header[index] = fmt.Sprintf("%s%s", t.header[index], DOWN_ARROW)

	return t
}

func (t *table) Print() {
	t.resetVariables()
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
	c.Println(t.buildRow(-1, t.header))
}

func (t *table) printRows() {
	for index, row := range t.rows {
		fmt.Println(t.buildRow(index, row))
	}
}

func (t *table) sortRows() {
	t.rows.SortBy(t.sortColumn)
}

func (t *table) buildRow(rowIndex int, row Row) (p string) {
	for columnIndex, columnValue := range row {
		p = p + t.buildColumn(rowIndex, columnIndex, fmt.Sprint(columnValue))
	}
	return p
}

func (t *table) buildColumn(rowIndex, columnIndex int, columnValue string) string {
	if rowIndex >= 0 && t.columns[columnIndex].Percentage == true {
		columnValue = fmt.Sprintf("%d%%", percentage(t.Min(columnIndex), t.Max(columnIndex), cast.StringToFloat64(columnValue)))
	}

	if rowIndex >= 0 && t.columns[columnIndex].ZeroFill == true {
		columnValue = t.printZeroFill(rowIndex, columnIndex, columnValue)
	}

	if t.columns[columnIndex].Alignment == Right {
		return t.lenOffset(columnValue, t.stats[columnIndex].Len) + columnValue + t.printPadding()
	}

	return columnValue + t.lenOffset(columnValue, t.stats[columnIndex].Len) + t.printPadding()
}

func (t *table) resetVariables() {
	t.filtered = 0
	t.width = 0
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

func (t *table) calculateColumnStats() {
	for index, value := range t.header {
		t.stats[index] = Stats{
			Len: t.calculateColumnLen(index, len(fmt.Sprint(value))),
			Min: t.calculateColumnMin(index, cast.InterfaceToFloat64(t.rows[0][index])),
			Max: t.calculateColumnMax(index, cast.InterfaceToFloat64(t.rows[0][index])),
			Sum: t.calculateColumnSum(index),
		}
	}
}

func (t *table) calculateColumnMin(index int, value float64) float64 {
	for x := 1; x < len(t.rows); x++ {
		value = math.Min(value, cast.InterfaceToFloat64(t.rows[x][index]))
	}

	return value
}

func (t *table) calculateColumnMax(index int, value float64) float64 {
	for x := 1; x < len(t.rows); x++ {
		value = math.Max(value, cast.InterfaceToFloat64(t.rows[x][index]))
	}

	return value
}

func (t *table) calculateColumnSum(index int) (sum float64) {
	for x := 0; x < len(t.rows); x++ {
		sum = sum + cast.InterfaceToFloat64(t.rows[x][index])
	}

	return sum
}

func (t *table) calculateColumnLen(index int, value int) int {
	for x := 0; x < len(t.rows); x++ {
		value = int(math.Max(float64(value), float64(len(fmt.Sprint(t.rows[x][index])))))
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

func (t *table) printZeroFill(rowIndex, columnIndex int, columnValue string) string {
	var m float64
	var n float64

	if v, err := strconv.ParseFloat(columnValue, 64); err == nil {
		n = v
		if v < 1e-6 {
			m = 1e6
		} else if v < 1e-3 {
			m = 1e3
		} else {
			m = 1
		}
	}
	return fmt.Sprintf(
		"%[1]*.[2]*[3]f",
		t.columns[columnIndex].Precision,
		t.columns[columnIndex].Scale, n*m)
}

func percentage(min, max, val float64) int {
	return int(((val-min)/(max-min))*99 + 1)
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
// Summary
// Limit
// Center: marginLeft, marginTop, marginBottom, marginRight
