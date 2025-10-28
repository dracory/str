package str

import "regexp"

// Start ensures the string begins with the provided prefix without duplicating it.
// If the prefix is empty, the original string is returned.
func Start(str, prefix string) string {
	if prefix == "" {
		return str
	}

	re := regexp.MustCompile("^(?:" + regexp.QuoteMeta(prefix) + ")+")
	trimmed := re.ReplaceAllString(str, "")
	return prefix + trimmed
}
