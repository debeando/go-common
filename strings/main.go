package strings

import (
	"regexp"
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

func ToCamel(s string) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := true

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

func addWordBoundariesToNumbers(s string) string {
	numberSequence := regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
	numberReplacement := []byte(`$1 $2 $3`)

	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}
