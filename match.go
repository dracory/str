package str

import "regexp"

// Match returns the first substring that matches the provided regular expression.
// If the pattern is empty, the original string is returned unchanged.
func Match(in, pattern string) string {
	if pattern == "" {
		return in
	}

	re := regexp.MustCompile(pattern)
	return re.FindString(in)
}
