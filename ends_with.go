package str

import "strings"

// EndsWith returns true if str ends with any of the provided values.
// Empty values are ignored. If no non-empty values are provided, returns false.
func EndsWith(str string, values ...string) bool {
	for _, value := range values {
		if value != "" && strings.HasSuffix(str, value) {
			return true
		}
	}

	return false
}
