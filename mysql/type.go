package mysql

import (
	"strconv"
	"strings"
)

func ParseNumberValue(value string) (int64, bool) {
	value = strings.ToLower(value)

	if value == "yes" || value == "on" {
		return 1, true
	}

	if value == "no" || value == "off" {
		return 0, true
	}

	if val, err := strconv.ParseInt(value, 10, 64); err == nil {
		return val, true
	}

	return 0, false
}

func ClearUser(u string) string {
	index := strings.Index(u, "[")
	if index > 0 {
		return u[0:index]
	}
	return u
}

func MaximumValueSigned(dataType string) uint64 {
	switch dataType {
	case "tinyint":
		return 127
	case "smallint":
		return 32767
	case "mediumint":
		return 8388607
	case "int":
		return 2147483647
	case "bigint":
		return 9223372036854775807
	}
	return 0
}

func MaximumValueUnsigned(dataType string) uint64 {
	switch dataType {
	case "tinyint":
		return 255
	case "smallint":
		return 65535
	case "mediumint":
		return 16777215
	case "int":
		return 4294967295
	case "bigint":
		return 18446744073709551615
	}
	return 0
}
