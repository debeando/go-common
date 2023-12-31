package digest

import (
	"strings"
	"unicode"
)

func Digest(s string) string {
	sql := []rune(strings.Trim(strings.ToLower(s), " "))
	comment := false
	endnumber := []rune{' ', ',', '+', '-', '*', '/', '^', '%', '(', ')'}
	length := len(sql)
	list := false
	multiline := false
	number := false
	quote := rune(0)
	result := []rune("")
	values := false
	whitespace := false
	whitespaces := 0

	IsNumber := func(r rune) bool {
		if unicode.IsNumber(r) || r == '.' {
			return true
		}
		return false
	}

	IsEndNumber := func(r rune) bool {
		for z := 0; z < len(endnumber); z++ {
			if r == endnumber[z] {
				return true
			}
		}
		return false
	}

	for x := 0; x < length; x++ {
		// Uncomment only for debug:
		// fmt.Printf("--> %02d/%d %s %s\n", x, length, string(sql[x]), string(result))

		// Remove comments:
		if !comment && !multiline && sql[x] == '#' {
			comment = true
		} else if comment && !multiline && sql[x] == '\n' {
			comment = false
			continue
		}

		if !comment && !multiline && sql[x] == '-' && x+1 < length && sql[x+1] == '-' {
			comment = true
		} else if comment && !multiline && sql[x] == '\n' {
			comment = false
			continue
		}

		if !comment && sql[x] == '/' && x+1 < length && sql[x+1] == '*' {
			comment = true
			multiline = true
		} else if comment && multiline && sql[x] == '*' && x+1 < length && sql[x+1] == '/' {
			x += 1
			comment = false
			multiline = false
			continue
		}

		if comment {
			continue
		}

		// Remove new lines:
		if sql[x] == '\n' || sql[x] == '\r' {
			sql[x] = ' '
			whitespace = true
			number = false
		}

		// Remove whitespaces until semicolon:
		if sql[x] == ' ' {
			for y := 1; y < (length - x); y++ {
				if whitespaces >= 0 && sql[x+y] == ';' {
					return string(append(result, ';'))
				} else if sql[x+y] == ' ' {
					whitespaces++
				} else {
					break
				}
			}
		}

		// Remove literals inside of list " IN (":
		if x >= 1 && sql[x-1] == ' ' && sql[x] == 'i' && x+1 < length && sql[x+1] == 'n' {
			for y := 2; y < (length - x); y++ {
				if sql[x+y] == '(' {
					list = true
					break
				} else if unicode.IsLetter(sql[x+y]) {
					break
				}
			}
		}

		if list {
			if !values && sql[x] == '(' {
				values = true
			} else if values && sql[x] == ')' {
				list = false
				values = false
				result = append(result, '?')
			} else if values {
				whitespace = false
				continue
			}
		}

		// Remove whitespaces:
		if quote == 0 && sql[x] == ' ' {
			whitespace = true
			number = false
		} else if quote == 0 && whitespace {
			whitespace = false
			result = append(result, ' ')
		}

		if whitespace {
			continue
		}

		// Remove backtick
		if quote == 0 && sql[x] == '`' {
			continue
		}

		// Remove string between quotes:
		if quote == 0 && (sql[x] == '"' || sql[x] == '\'') {
			quote = sql[x]
			result = append(result, '\'')
		} else if quote > 0 && sql[x] == '\\' && x+1 < length && sql[x+1] == quote {
			x += 1
		} else if sql[x] == quote {
			quote = 0
			result = append(result, '?')
			result = append(result, '\'')
			continue
		}

		if quote > 0 {
			continue
		}

		// Remove numbers:
		if !number && IsNumber(sql[x]) {
			number = true

			// Check to skip word composed with number and letter:
			for y := x; y >= 0; y-- {
				if IsEndNumber(sql[y]) {
					break
				} else {
					if unicode.IsLetter(sql[y]) {
						number = false
					}
				}
			}

			for y := 0; y < (length - x); y++ {
				if IsEndNumber(sql[x+y]) {
					break
				} else {
					if unicode.IsLetter(sql[x+y]) {
						number = false
					}
				}
			}

			// Add ? symbol to remove nombre:
			if number {
				result = append(result, '?')
			}
		}

		if number && !IsNumber(sql[x]) {
			number = false
		}

		if number {
			continue
		}

		result = append(result, sql[x])
	}

	return string(result)
}
