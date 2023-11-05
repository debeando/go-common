package digest

import (
	"sort"

	"github.com/debeando/go-common/mysql/sql/parser/slow"
)

type List []Stats

var FilterByQueryTime float64

func (l List) Index(q slow.Query) (int, bool) {
	for i := range l {
		if l[i].ID == q.DigestID {
			return i, true
		}
	}
	return 0, false
}

func (l *List) Add(q slow.Query) {
	index, exists := l.Index(q)

	if !exists {
		stats := Stats{}
		stats.Append(q)
		(*l) = append((*l), stats)
	} else {
		(*l)[index].Append(q)
	}
}

func (l List) Count() (sum int) {
	for i := range l {
		sum = sum + l[i].Count
	}

	return sum
}

func (l List) FilterByQueryTime(v float64) {
	FilterByQueryTime = v
}

func (l List) Filtered() (sum int) {
	for i := range l {
		if l[i].Time.Query < FilterByQueryTime {
			sum = sum + l[i].Count
		}
	}

	return sum
}

func (l List) Analyzed() (sum int) {
	for i := range l {
		if l[i].Time.Query >= FilterByQueryTime {
			sum = sum + l[i].Count
		}
	}

	return sum
}

func (l List) Unique() (sum int) {
	for i := range l {
		if l[i].Time.Query >= FilterByQueryTime {
			sum++
		}
	}

	return sum
}

func (l List) ScoreMax() (max float64) {
	for i := range l {
		if l[i].Time.Query >= FilterByQueryTime {
			if l[i].Score > max {
				max = l[i].Score
			}
		}
	}

	return max
}

func (l List) ScoreMin() (min float64) {
	for i := range l {
		if l[i].Time.Query >= FilterByQueryTime {
			if l[i].Score < min {
				min = l[i].Score
			}
		}
	}

	return min
}

func (l *List) Clean() {
	for i := len((*l)) - 1; i >= 0; i-- {
		if (*l)[i].Time.Query < FilterByQueryTime {
			(*l).Remove(i)
		}
	}
}

func (l *List) Remove(i int) {
	if len((*l)) > i {
		(*l) = append((*l)[:i], (*l)[i+1:]...)
	}
}

// Len is part of sort.Interface.
func (l List) Len() int {
	return len(l)
}

// Less is part of sort.Interface.
// We use count as the value to sort by
func (l List) Less(i, j int) bool {
	return l[i].Score > l[j].Score
}

// Swap is part of sort.Interface.
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) SortByScore() {
	sort.Sort(l)
}
