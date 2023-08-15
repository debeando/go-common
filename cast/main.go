package cast

import (
	"strconv"
	"strings"
)

func StringToInt(value string) int {
	i, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0
	}
	return i
}
