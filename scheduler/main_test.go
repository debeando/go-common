package scheduler_test

import (
	mock "time"

	"testing"

	"github.com/debeando/go-common/scheduler"
	"github.com/debeando/go-common/time"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	s := scheduler.Scheduler{
		Start: "07:00",
		End:   "20:00",
	}

	start, _ := time.StringToTime(s.Start)
	end, _ := time.StringToTime(s.End)

	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 06, 59, 00, 000000000, mock.UTC)
	}
	assert.False(t, time.BetweenNow(start, end))

	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 07, 00, 00, 000000000, mock.UTC)
	}
	assert.False(t, time.BetweenNow(start, end))

	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 07, 01, 00, 000000000, mock.UTC)
	}
	assert.True(t, time.BetweenNow(start, end))

	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 18, 45, 00, 000000000, mock.UTC)
	}
	assert.True(t, time.BetweenNow(start, end))

	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 19, 59, 00, 000000000, mock.UTC)
	}
	assert.True(t, time.BetweenNow(start, end))

	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 20, 00, 00, 000000000, mock.UTC)
	}
	assert.False(t, time.BetweenNow(start, end))

	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 20, 01, 00, 000000000, mock.UTC)
	}
	assert.False(t, time.BetweenNow(start, end))
}

func TestDelta(t *testing.T) {
	time.Now = func() mock.Time {
		return mock.Date(0, 1, 1, 20, 01, 40, 000000000, mock.UTC)
	}

	assert.Equal(t, scheduler.DeltaSeconds(), mock.Duration(20))
}
