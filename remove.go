package str

import "strings"

// Remove deletes all occurrences of the provided values from the input string.
func Remove(str string, values ...string) string {
	if len(values) == 0 {
		return str
	}

	result := str
	for _, value := range values {
		result = strings.ReplaceAll(result, value, "")
	}

	return result
}
