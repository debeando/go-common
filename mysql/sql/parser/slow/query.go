package slow

import (
	"strings"
	"time"

	"github.com/debeando/go-common/cast"
	"github.com/debeando/go-common/crypt"
	"github.com/debeando/go-common/mysql/sql/parser/digest"
)

type Query struct {
	Time         time.Time
	Timestamp    int64
	User         string
	ID           int64
	QueryTime    float64
	LockTime     float64
	RowsSent     int64
	RowsExamined int64
	Raw          string
	Digest       string
	DigestID     string
}

func QueryParser(query string) Query {
	property := map[string]string{}
	whiteSpaceStart := 0
	whiteSpaceEnd := 0
	startQuery := 0

	p := []rune(query)
	l := len(p)

	for x := 0; x < l; x++ {
		// Register first White Space:
		if p[x] == ' ' {
			whiteSpaceStart = x
		}

		// Start second loop to find next property:
		if p[x] == ':' && p[x+1] == ' ' {
			for y := x + 1; y < l; y++ {
				// Stop when is finished header and start SQL:
				if p[y] == '\n' && p[y+1] != '#' {
					whiteSpaceEnd = y
					break
				}

				// Remove header comments:
				if p[y] == '#' || p[y] == '\n' || p[y] == '\r' {
					p[y] = ' '
				}

				// Register last White Space:
				if p[y] == ' ' {
					whiteSpaceEnd = y
					continue
				}

				// Stop when find next property:
				if p[y] == ':' && p[y+1] == ' ' {
					break
				}
			}

			key := string(p[whiteSpaceStart:x])
			key = strings.TrimSpace(key)
			key = strings.ToLower(key)

			value := strings.TrimSpace(string(p[x+1 : whiteSpaceEnd]))

			property[key] = value
		}

		// Find timestamp value:
		if (x+24) <= l && string(p[x:x+14]) == "SET timestamp=" {
			property["timestamp"] = string(p[x+14 : x+24])
			startQuery = x + 25
		}
	}

	property["query"] = string(p[startQuery:l])
	property["query"] = strings.Trim(property["query"], "\n")
	property["user"] = UserParser(property["user@host"])

	digestQuery := digest.Digest(property["query"])

	return Query{
		Time:         cast.StringToDateTime(property["time"], "2006-01-02T15:04:05.000000Z"),
		ID:           cast.StringToInt64(property["id"]),
		LockTime:     cast.StringToFloat64(property["lock_time"]),
		QueryTime:    cast.StringToFloat64(property["query_time"]),
		RowsExamined: cast.StringToInt64(property["rows_examined"]),
		RowsSent:     cast.StringToInt64(property["rows_sent"]),
		Timestamp:    cast.StringToInt64(property["timestamp"]),
		Raw:          property["query"],
		User:         property["user"],
		Digest:       digestQuery,
		DigestID:     crypt.MD5(digestQuery),
	}
}

func UserParser(u string) string {
	p := []rune(u)

	for x := 0; x < len(p); x++ {
		if p[x] == '[' {
			return string(p[0:x])
		}
	}

	return ""
}
