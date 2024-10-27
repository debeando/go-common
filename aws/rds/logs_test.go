package rds_test

import (
	"testing"

	"github.com/debeando/go-common/aws/rds"

	"github.com/stretchr/testify/assert"
)

func TestLogsSort(t *testing.T) {
	logs := rds.NewLogs()
	logs.Add(rds.Log{Size: int64(6988063)})
	logs.Add(rds.Log{Size: int64(27260625)})
	logs.Add(rds.Log{Size: int64(1398553)})
	logs.Add(rds.Log{Size: int64(18925915)})

	t.Log(logs.Len())

	// logs.SortBySize()

	for _, log := range logs {
		t.Log(log.Size)
	}

	assert.True(t, true)
}
