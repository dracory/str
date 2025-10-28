package str

import (
	"strings"
	"unicode"
)

// RTrim removes trailing characters from the input string.
// When no characters are provided, Unicode whitespace is trimmed.
func RTrim(str string, characters ...string) string {
	if len(characters) == 0 || characters[0] == "" {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}

	return strings.TrimRight(str, characters[0])
}
