package str

import "strings"

// ReplaceFirst replaces the first occurrence of search with replace.
// If search is empty, the original string is returned unchanged.
func ReplaceFirst(str, search, replace string) string {
	if search == "" {
		return str
	}

	return strings.Replace(str, search, replace, 1)
}
