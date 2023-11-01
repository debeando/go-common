package slow

import (
	"strings"
)

func LogsParser(in <-chan string, out chan<- string) {
	var buffer string
	var isHeader bool
	var isQuery bool

	for line := range in {
		e := ""
		l := len(line)

		if isQuery == false && strings.HasPrefix(line, "# ") {
			isHeader = true
		}

		if isHeader == true && l >= 6 {
			buffer += line + "\n"

			s := string(line[0:6])
			s = strings.ToUpper(s)

			if s == "SELECT" || s == "INSERT" || s == "UPDATE" || s == "DELETE" {
				isQuery = true
			}
		}

		if l > 1 {
			e = string(line[l-1:])
		} else {
			e = string(line)
		}

		if isQuery == true && e == ";" {
			out <- strings.TrimRight(buffer, "\n")

			buffer = ""
			isHeader = false
			isQuery = false
		}
	}
}
