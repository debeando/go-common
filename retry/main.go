package retry

import (
	"errors"
	"time"
)

func Do(attempts uint8, sleep time.Duration, fn func() bool) error {
	for attempt := uint8(0); attempt < attempts; attempt++ {
		time.Sleep(sleep)

		if fn() {
			return nil
		}
	}
	return errors.New("Exhausted attempts")
}
