package str

import "strings"

// Headline converts a string into headline case.
// When the input already contains whitespace, it is simply title-cased.
// Otherwise, the string is normalized, converted to studly case, split on
// uppercase boundaries, and the resulting words are joined with spaces.
func Headline(in string) string {
	if strings.Contains(in, " ") {
		return Title(in)
	}

	normalized := strings.ReplaceAll(in, "-", "_")
	normalized = strings.ReplaceAll(normalized, " ", "_")

	studly := Studly(normalized)
	words := splitHeadlineWords(studly)
	return strings.Join(words, " ")
}

func splitHeadlineWords(s string) []string {
	if s == "" {
		return nil
	}

	runes := []rune(s)
	var words []string
	start := 0

	for i := 1; i < len(runes); i++ {
		current := runes[i]
		prev := runes[i-1]

		if isBoundary(prev, current, runes, i) {
			words = append(words, string(runes[start:i]))
			start = i
		}
	}

	words = append(words, string(runes[start:]))
	return words
}

func isBoundary(prev, current rune, runes []rune, index int) bool {
	if isLower(prev) && isUpper(current) {
		return true
	}
	if isUpper(prev) && isUpper(current) && index+1 < len(runes) && isLower(runes[index+1]) {
		return true
	}
	if isLetter(prev) && !isLetter(current) {
		return true
	}
	if !isLetter(prev) && isLetter(current) {
		return true
	}
	return false
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isLetter(r rune) bool {
	return isUpper(r) || isLower(r)
}
