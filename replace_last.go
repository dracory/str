package str

import "strings"

// ReplaceLast replaces the last occurrence of search with replace.
// If search is empty or not found, the original string is returned unchanged.
func ReplaceLast(str, search, replace string) string {
	if search == "" {
		return str
	}

	index := strings.LastIndex(str, search)
	if index == -1 {
		return str
	}

	return str[:index] + replace + str[index+len(search):]
}
