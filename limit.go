package str

import "unicode/utf8"

// Limit truncates the string to the given length, appending the provided ending
// (or "..." by default) when truncation occurs. Length is measured in runes.
func Limit(str string, limit int, end ...string) string {
	if limit < 0 {
		return str
	}

	if utf8.RuneCountInString(str) <= limit {
		return str
	}

	suffix := "..."
	if len(end) > 0 && end[0] != "" {
		suffix = end[0]
	}

	if limit == 0 {
		return suffix
	}

	runes := []rune(str)
	return string(runes[:limit]) + suffix
}
