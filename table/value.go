package table

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"github.com/debeando/go-common/cast"
)

type Field struct {
	Value any
}

func (f Field) IsString() bool {
	return reflect.ValueOf(f.Value).Kind() == reflect.String
}

func (f Field) ToString() string {
	return fmt.Sprintf("%v", f.Value)
}

func (f Field) ToFloat64() float64 {
	return cast.InterfaceToFloat64(f.Value)
}

func (f Field) Len() int {
	if f.IsString() {
		return len(f.ToString())
	}
	return 0
}

func (f *Field) Truncate(t int) {
	if f.IsString() && t < f.Len() && t > 0 {
		(*f).Value = f.ToString()[:t]
	}
}

func (f *Field) Clear() {
	if f.IsString() {
		t := f.ToString()
		t = strings.TrimSpace(t)
		t = strings.ReplaceAll(t, "\n", " ")
		t = strings.ReplaceAll(t, "\r", " ")
		t = strings.ReplaceAll(t, "  ", " ")

		(*f).Value = t
	}
}

func (f Field) ZeroFill(precision, scale int) string {
	var m float64
	var n float64

	if v, err := strconv.ParseFloat(f.ToString(), 64); err == nil {
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
		precision,
		scale, n*m)
}

func (f Field) Percentage(min, max float64) int {
	return int(((f.ToFloat64()-min)/(max-min))*99 + 1)
}

func (f Field) EvalCondition(condition string) bool {
	fs := token.NewFileSet()
	tv, _ := types.Eval(
		fs,
		nil,
		token.NoPos,
		fmt.Sprintf("%s %s", f.ToString(), condition))

	return constant.BoolVal(tv.Value)
}
