package array

func StringIn(key string, list []string) bool {
	for _, l := range list {
		if l == key {
			return true
		}
	}
	return false
}
