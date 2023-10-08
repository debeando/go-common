package maps

import (
	"sort"
)

func Keys(v []map[string]string) (keys []string) {
	if len(v) > 0 {
		for key := range v[0] {
			keys = append(keys, key)
		}
		sort.Strings(keys)
	}
	return
}

func In(key string, list map[string]string) bool {
	if _, ok := list[key]; ok {
		return true
	}
	return false
}

func ComparteString(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}
