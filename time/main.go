package time

import (
	"time"
)

var Now func() time.Time

func init() {
	Now = func() time.Time {
		return time.Now()
	}
}

func StringToTime(in string) (time.Time, error) {
	return time.Parse("15:04", in)
}

func NowUTCf() string {
	return Now().UTC().Format("2006-01-02 15:04:05")
}

func BetweenNow(b time.Time, a time.Time) bool {
	n := Now().UTC()

	return n.After(b) && n.Before(a)
}
