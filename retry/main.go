package retry

import (
	"errors"
	"time"
)

func Do(attempts int, sleep time.Duration, fn func() bool) error {
	for attempt := 0; attempt < attempts; attempt++ {
		time.Sleep(sleep * time.Second)

		if fn() {
			return nil
		}
	}
	return errors.New("Exhausted attempts")
}
