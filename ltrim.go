package str

import (
	"strings"
	"unicode"
)

// LTrim removes leading characters from the input string.
// When no character set is provided, leading Unicode whitespace is trimmed.
func LTrim(in string, characters ...string) string {
	if len(characters) == 0 || characters[0] == "" {
		return strings.TrimLeftFunc(in, unicode.IsSpace)
	}

	return strings.TrimLeft(in, characters[0])
}
