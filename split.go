package str

import "regexp"

// Split divides the string using the provided regex pattern. The optional limit
// matches regexp.Split behaviour: default -1 keeps all substrings.
func Split(str, pattern string, limit ...int) []string {
	if pattern == "" {
		return []string{str}
	}

	re := regexp.MustCompile(pattern)
	splitLimit := -1
	if len(limit) > 0 {
		splitLimit = limit[0]
	}

	return re.Split(str, splitLimit)
}
