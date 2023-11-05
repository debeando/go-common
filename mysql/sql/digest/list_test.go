package digest_test

import (
	"testing"

	"github.com/debeando/go-common/mysql/sql/digest"
	"github.com/debeando/go-common/mysql/sql/parser/slow"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	queries := &digest.List{
		digest.Stats{ID: "a"},
		digest.Stats{ID: "b"},
		digest.Stats{ID: "c"},
	}

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
		i, e := queries.Index(test.Query)
		assert.Equal(t, i, test.Index)
		assert.Equal(t, e, test.Exist)
	}
}

func TestAddAndCount(t *testing.T) {
	queries := &digest.List{}
	queries.Add(slow.Query{DigestID: "a"})
	queries.Add(slow.Query{DigestID: "a"})
	queries.Add(slow.Query{DigestID: "b"})

	assert.Equal(t, queries.Count(), 3)
}

func TestAddAndUnique(t *testing.T) {
	queries := &digest.List{}
	queries.Add(slow.Query{DigestID: "a"})
	queries.Add(slow.Query{DigestID: "a"})
	queries.Add(slow.Query{DigestID: "b"})

	assert.Equal(t, queries.Unique(), 2)
}
