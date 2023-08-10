package time_test

import (
	mock "time"

	"testing"

	"github.com/debeando/go-common/time"
	"github.com/stretchr/testify/assert"
)

func TestStringToTime(t *testing.T) {
	x, e := time.StringToTime("12:59")
	assert.NoError(t, e)
	assert.Equal(t, x, mock.Date(0000, 01, 01, 12, 59, 00, 000000000, mock.UTC))

	x, e = time.StringToTime("12:61")
	assert.Error(t, e)
	assert.Empty(t, x)
}

func TestNowUTCf(t *testing.T) {
	time.Now = func() mock.Time {
		return mock.Date(1980, 05, 19, 18, 45, 00, 000000000, mock.UTC)
	}

	assert.Equal(t, time.NowUTCf(), "1980-05-19 18:45:00")
}

func TestBetweenNow(t *testing.T) {
	time.Now = func() mock.Time {
		return mock.Date(0000, 01, 01, 13, 30, 00, 000000000, mock.UTC)
	}

	assert.True(t, time.BetweenNow(
		mock.Date(0000, 01, 01, 13, 00, 00, 000000000, mock.UTC), // Before (Antes)
		mock.Date(0000, 01, 01, 13, 59, 00, 000000000, mock.UTC), // After  (Después)
	))
	assert.False(t, time.BetweenNow(
		mock.Date(0000, 01, 01, 11, 00, 00, 000000000, mock.UTC), // Before (Antes)
		mock.Date(0000, 01, 01, 11, 59, 00, 000000000, mock.UTC), // After  (Después)
	))
	assert.False(t, time.BetweenNow(
		mock.Date(0000, 01, 01, 20, 00, 00, 000000000, mock.UTC), // Before (Antes)
		mock.Date(0000, 01, 01, 20, 59, 00, 000000000, mock.UTC), // After  (Después)
	))
}
