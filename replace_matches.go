package str

import "regexp"

// ReplaceMatches replaces all substrings matching the pattern with the
// provided replacement string. If the pattern is empty, the original string is
// returned unchanged.
func ReplaceMatches(str, pattern, replace string) string {
	if pattern == "" {
		return str
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return str
	}

	return re.ReplaceAllString(str, replace)
}
