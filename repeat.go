package str

import "strings"

// Repeat repeats the input string the specified number of times.
func Repeat(in string, times int) string {
	if times <= 0 {
		return ""
	}

	return strings.Repeat(in, times)
}
