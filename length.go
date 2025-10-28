package str

import "unicode/utf8"

// Length returns the number of runes in the given string.
func Length(in string) int {
	return utf8.RuneCountInString(in)
}
