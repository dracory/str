package str

import "unicode"

// Squish collapses consecutive whitespace into a single space and trims
// leading/trailing whitespace.
func Squish(str string) string {
	var result []rune
	spacePending := false

	for _, r := range str {
		if unicode.IsSpace(r) {
			spacePending = len(result) > 0
			continue
		}

		if spacePending {
			result = append(result, ' ')
			spacePending = false
		}

		result = append(result, r)
	}

	return string(result)
}
