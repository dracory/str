package str

import (
	"strings"
	"unicode/utf8"
)

// Mask replaces a portion of the string with a repeated masking character.
// The index marks the starting rune offset, with negative values counting
// backwards from the end. If length is omitted the remainder of the string is
// masked. When the masking character is empty or the target segment is empty,
// the original string is returned unchanged.
func Mask(in, character string, index int, length ...int) string {
	if character == "" {
		return in
	}

	segment := Substr(in, index, length...)
	if segment == "" {
		return in
	}

	strLen := utf8.RuneCountInString(in)
	startIndex := index
	if index < 0 {
		if index < -strLen {
			startIndex = 0
		} else {
			startIndex = strLen + index
		}
	}

	start := Substr(in, 0, startIndex)
	segmentLen := utf8.RuneCountInString(segment)
	end := Substr(in, startIndex+segmentLen)

	maskRune := Substr(character, 0, 1)
	return start + strings.Repeat(maskRune, segmentLen) + end
}
