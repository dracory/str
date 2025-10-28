package str

import "strings"

// ChopStart removes the first matching prefix found in needle or more from the start of str.
// If none of the provided prefixes match, the original string is returned unchanged.
func ChopStart(str string, needle string, more ...string) string {
	candidates := append([]string{needle}, more...)

	for _, v := range candidates {
		if v == "" {
			continue
		}
		if strings.HasPrefix(str, v) {
			return strings.TrimPrefix(str, v)
		}
	}

	return str
}
