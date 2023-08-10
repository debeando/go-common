package time

import (
	"time"
)

var Now func() time.Time

func init() {
	Now = func() time.Time {
		return time.Now().UTC()
	}
}

func StringToTime(in string) (time.Time, error) {
	return time.Parse("15:04", in)
}

func NowUTCf() string {
	return Now().Format("2006-01-02 15:04:05")
}

func BetweenNow(b time.Time, a time.Time) bool {
	n := Now()

	// Truncate year, month and day.
	n = n.AddDate(-int(n.Year())+1, -int(n.Month())+1, -int(n.Day())+1)
	b = b.AddDate(-int(b.Year())+1, -int(b.Month())+1, -int(b.Day())+1)
	a = a.AddDate(-int(a.Year())+1, -int(a.Month())+1, -int(a.Day())+1)

	return n.After(b) && n.Before(a)
}
