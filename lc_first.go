package str

import "unicode"

// LcFirst lowercases the first rune of the string.
func LcFirst(in string) string {
	if in == "" {
		return ""
	}

	runes := []rune(in)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
