package slow_test

import (
	"testing"
	"time"

	"github.com/debeando/go-common/mysql/sql/parser/slow"

	"github.com/stretchr/testify/assert"
)

var logs = []string{
	"# Time: 2023-10-27T08:51:32.366670Z",
	"# User@Host: test[test] @  [127.0.0.1]  Id:   725",
	"# Query_time: 0.000135  Lock_time: 0.000002 Rows_sent: 11  Rows_examined: 11",
	"SET timestamp=1698396692;",
	"SELECT * FROM foo WHERE deleted_at IS NULL;",
	"# Time: 2023-10-27T08:51:34.376283Z",
	"# User@Host: test[test] @ [127.0.0.1]  Id:  3178",
	"# Query_time: 0.019560  Lock_time: 0.000002 Rows_sent: 1  Rows_examined: 56914",
	"SET timestamp=1698397201;",
	"SELECT",
	" count(*)",
	"",
	"FROM foo",
	";",
	"# Time: 2023-10-27T08:51:34.376283Z",
	"# User@Host: test[test] @  [127.0.0.1]  Id:  1303",
	"# Query_time: 0.000126  Lock_time: 0.000002 Rows_sent: 18  Rows_examined: 3583",
	"SET timestamp=1698396694;",
	"SELECT *",
	"      FROM foo;",
}

func TestLogsParser(t *testing.T) {
	logsCount := 0
	channelIn := make(chan string)
	channelOut := make(chan string)

	defer close(channelIn)

	go slow.LogsParser(channelIn, channelOut)

	go func() {
		defer close(channelOut)
		for query := range channelOut {
			t.Log("\n", query)
			logsCount++
		}
	}()

	for _, line := range logs {
		channelIn <- line
	}

	time.Sleep(1 * time.Second)

	assert.Equal(t, 3, logsCount)
}
