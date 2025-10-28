package str

import "strings"

// ContainsAll returns true if str contains all the provided values.
// Empty values are treated as matches (consistent with strings.Contains behaviour).
func ContainsAll(str string, values ...string) bool {
	for _, value := range values {
		if !strings.Contains(str, value) {
			return false
		}
	}

	return true
}
