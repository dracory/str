package str

import "strings"

// ChopEnd removes the first matching suffix found in needle or more from the end of str.
// If none of the provided suffixes match, the original string is returned unchanged.
func ChopEnd(str string, needle string, more ...string) string {
	candidates := append([]string{needle}, more...)

	for _, v := range candidates {
		if v == "" {
			continue
		}
		if strings.HasSuffix(str, v) {
			return strings.TrimSuffix(str, v)
		}
	}

	return str
}
