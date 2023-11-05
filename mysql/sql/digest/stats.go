package digest

import (
	"github.com/debeando/go-common/mysql/sql/parser/slow"
)

type Stats struct {
	ID    string
	Count int
	Score float64
	Time  struct {
		Query float64
		Lock  float64
	}
	Rows struct {
		Sent     int64
		Examined int64
	}
	Queries []slow.Query
}

func (s *Stats) Append(q slow.Query) {
	s.ID = q.DigestID
	s.Queries = append(s.Queries, q)
	s.Count = len(s.Queries)

	s.SetQueryTime(q.QueryTime)
	s.SetLockTime(q.LockTime)
	s.SetRowsSent(q.RowsSent)
	s.SetRowsExamined(q.RowsExamined)
	s.SetScore()
}

func (s *Stats) SetQueryTime(t float64) {
	if t > s.Time.Query {
		s.Time.Query = t
	}
}

func (s *Stats) SetLockTime(t float64) {
	if t > s.Time.Lock {
		s.Time.Lock = t
	}
}

func (s *Stats) SetRowsSent(t int64) {
	if t > s.Rows.Sent {
		s.Rows.Sent = t
	}
}

func (s *Stats) SetRowsExamined(t int64) {
	if t > s.Rows.Examined {
		s.Rows.Examined = t
	}
}

func (s *Stats) SetScore() {
	s.Score = (s.Time.Query * 0.4) +
		(s.Time.Lock * 0.1) +
		(float64(s.Count) * 0.2) +
		(float64(s.Rows.Examined) * 0.3)
}
