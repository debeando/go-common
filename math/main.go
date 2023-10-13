package math

func Percentage(value int64, max uint64) float64 {
	v := float64(value)
	m := float64(max)
	if v >= 0 && m > 0 {
		return (v / m) * 100
	}
	return 0
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
