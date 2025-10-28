package str

import "strings"

// StartsWith returns true if the string starts with any of the provided values.
// When no values are supplied, it returns false.
func StartsWith(str string, values ...string) bool {
	if len(values) == 0 {
		return false
	}

	for _, value := range values {
		if strings.HasPrefix(str, value) {
			return true
		}
	}

	return false
}
