package str

import (
	"regexp"
	"strings"
)

// Replace substitutes occurrences of search with replace. Case sensitivity can
// be controlled via the optional caseSensitive flag (defaults to true).
func Replace(str, search, replace string, caseSensitive ...bool) string {
	if search == "" {
		return str
	}

	sensitive := true
	if len(caseSensitive) > 0 {
		sensitive = caseSensitive[0]
	}

	if !sensitive {
		pattern := regexp.MustCompile("(?i)" + regexp.QuoteMeta(search))
		return pattern.ReplaceAllString(str, replace)
	}

	return strings.ReplaceAll(str, search, replace)
}
