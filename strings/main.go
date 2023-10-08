package strings

import (
	"strings"
)

func SplitKeyAndValue(s *string) (key string, value string) {
	kv := strings.SplitN(*s, "=", 2)
	if len(kv) == 2 {
		return strings.TrimSpace(strings.ToLower(kv[0])), kv[1]
	}
	return "", ""
}

func Trim(value *string) string {
	*value = strings.TrimSpace(*value)
	*value = strings.TrimRight(*value, "\"")
	*value = strings.TrimLeft(*value, "\"")
	return *value
}

func Escape(text string) string {
	return strings.Replace(text, "'", `\'`, -1)
}
