package utils

// InList 判断key是否在列表中
func InList(key string, List []string) bool {
	for _, s := range List {
		if key == s {
			return true
		}
	}
	return false
}
