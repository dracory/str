package str

import "regexp"

// MatchAll returns all substrings that match the provided regular expression.
// If the pattern is empty, the entire string is returned as the single result.
func MatchAll(in, pattern string) []string {
	if pattern == "" {
		return []string{in}
	}

	re := regexp.MustCompile(pattern)
	return re.FindAllString(in, -1)
}
