package str

import "strings"

// Trim removes the provided characters from both ends of the string. When no
// characters are provided, Unicode spaces are removed.
func Trim(str string, characters ...string) string {
	if len(characters) == 0 || characters[0] == "" {
		return strings.TrimSpace(str)
	}

	return strings.Trim(str, characters[0])
}
