package cast

import (
	"strconv"
	"strings"
	"time"
)

func StringToInt(value string) int {
	i, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0
	}
	return i
}

func StringToInt64(value string) int64 {
	i, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func StringToFloat64(value string) float64 {
	i, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return 0
	}
	return i
}

func ToDateTime(timestamp string, layout string) string {
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func InterfaceToInt64(value interface{}) int64 {
	if v, ok := value.(int64); ok {
		return v
	}
	return 0
}

func InterfaceToFloat64(value interface{}) float64 {
	if v, ok := value.(float64); ok {
		return float64(v)
	}
	return 0
}
