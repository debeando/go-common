package digest_test

import (
	"testing"

	"github.com/debeando/go-common/mysql/sql/digest"
	"github.com/debeando/go-common/mysql/sql/parser/slow"
	"github.com/stretchr/testify/assert"
)

var test_queries digest.List

// DIGEST ID                         SCORE  COUNT  Q. TIME  L. TIME  R. SENT  R. EXAMINED
// 78e8b4723e5360a2687c52a460da67ce  100    1      24.112   00.002   0        1725134
// 3b16991e6c95a21d59df0f441ef7251c  046    25     53.160   00.004   20       797170

func init() {
	test_queries.Add(slow.Query{DigestID: "a", QueryTime: 1.0, LockTime: 0.2, RowsExamined: 10000})
	test_queries.Add(slow.Query{DigestID: "b", QueryTime: 4.0, LockTime: 0.1, RowsExamined: 100000})
	test_queries.Add(slow.Query{DigestID: "b", QueryTime: 0.1, LockTime: 0.0, RowsExamined: 1000})
	test_queries.Add(slow.Query{DigestID: "c", QueryTime: 0.4, LockTime: 0.1, RowsExamined: 10000})
	test_queries.Add(slow.Query{DigestID: "d", QueryTime: 0.1, LockTime: 0.1, RowsExamined: 100})
	test_queries.Add(slow.Query{DigestID: "b", QueryTime: 1.3, LockTime: 0.0, RowsExamined: 10000})
	test_queries.Add(slow.Query{DigestID: "a", QueryTime: 0.1, LockTime: 0.2, RowsExamined: 1000})
	test_queries.Add(slow.Query{DigestID: "e", QueryTime: 0.0, LockTime: 0.0, RowsExamined: 10})
	test_queries.Add(slow.Query{DigestID: "b", QueryTime: 0.2, LockTime: 0.1, RowsExamined: 1000})
	test_queries.Add(slow.Query{DigestID: "c", QueryTime: 1.0, LockTime: 0.0, RowsExamined: 100000})
}

func TestIndex(t *testing.T) {
	var cases = []struct {
		Query slow.Query
		Index int
		Exist bool
	}{
		{
			Query: slow.Query{DigestID: "a"},
			Index: 0,
			Exist: true,
		},
		{
			Query: slow.Query{DigestID: "b"},
			Index: 1,
			Exist: true,
		},
		{
			Query: slow.Query{DigestID: "z"},
			Index: 0,
			Exist: false,
		},
	}

	for _, test := range cases {
		i, e := test_queries.Index(test.Query)
		assert.Equal(t, i, test.Index)
		assert.Equal(t, e, test.Exist)
	}
}

// func TestAddAndCount(t *testing.T) {
// 	// queries := &digest.List{}
// 	// queries.Add(slow.Query{DigestID: "a"})
// 	// queries.Add(slow.Query{DigestID: "a"})
// 	// queries.Add(slow.Query{DigestID: "b"})

// 	assert.Equal(t, test_queries.Count(), 3)
// }

// func TestClean(t *testing.T) {
// 	test_queries.FilterByQueryTime(3.0)
// 	test_queries.Clean()

// 	// assert.Equal(t, 1, test_queries.Count())

// 	t.Log(test_queries.Analyzed(), test_queries.Unique(), test_queries.ScoreMax(), test_queries.ScoreMin(), test_queries.Count(), test_queries.Len())
// }

// func TestAddAndUnique(t *testing.T) {
// 	// queries := &digest.List{}
// 	// queries.Add(slow.Query{DigestID: "a"})
// 	// queries.Add(slow.Query{DigestID: "a"})
// 	// queries.Add(slow.Query{DigestID: "b"})
// 	test_queries.FilterByQueryTime(0.0)

// 	assert.Equal(t, test_queries.Unique(), 2)
// }
