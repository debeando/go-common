package slow_test

import (
	"testing"
	"time"

	"github.com/debeando/go-common/mysql/sql/parser/slow"

	"github.com/stretchr/testify/assert"
)

func TestQueryParser(t *testing.T) {
	log := `
  # Time: 2023-10-27T08:51:34.376283Z
  # User@Host: test[test] @  [127.0.0.1]  Id:  1303
  # Query_time: 0.000126  Lock_time: 0.000002 Rows_sent: 18  Rows_examined: 3583
  SET timestamp=1698396694;
  SELECT *
        FROM foo ;`

	query := slow.QueryParser(log)

	expectedTime := time.Date(2023, 10, 27, 8, 51, 34, 376283, time.UTC)
	expectedTime = expectedTime.Round(time.Millisecond)
	actualTime := query.Time
	actualTime = actualTime.Add(-time.Duration(actualTime.Nanosecond()))

	assert.Equal(t, expectedTime, actualTime)
	assert.Equal(t, int64(1698396694), query.Timestamp)
	assert.Equal(t, int64(18), query.RowsSent)
	assert.Equal(t, int64(3583), query.RowsExamined)
	assert.Equal(t, int64(1303), query.ID)
	assert.Equal(t, float64(0.000126), query.QueryTime)
	assert.Equal(t, float64(0.000002), query.LockTime)
	assert.Equal(t, "select * from foo;", query.Digest)
	assert.Equal(t, "test", query.User)
}
