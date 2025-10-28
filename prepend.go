package str

import "strings"

// Prepend concatenates the provided values before the input string.
func Prepend(str string, values ...string) string {
	if len(values) == 0 {
		return str
	}

	return strings.Join(values, "") + str
}
