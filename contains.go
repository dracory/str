package str

import "strings"

// Contains returns true if str contains any of the provided values.
// Empty values are ignored. If no non-empty values are provided, returns false.
func Contains(str string, values ...string) bool {
	for _, value := range values {
		if value != "" && strings.Contains(str, value) {
			return true
		}
	}

	return false
}
