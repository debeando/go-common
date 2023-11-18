package table

import (
	"math"
)

type Column struct {
	Alignment  Alignment
	Name       string
	Percentage bool
	Precision  int
	Scale      int
	Truncate   int
	Width      int
	ZeroFill   bool
}

func (c *Column) SetName(s string) {
	(*c).Name = s
}

func (c Column) GetWidth() int {
	return int(math.Max(float64(len(c.Name)), float64(c.Width)))
}
