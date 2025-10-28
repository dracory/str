package str

import "strings"

// ExcerptOption configures the behaviour of Excerpt.
type ExcerptOption struct {
	Radius   int
	Omission string
}

// Excerpt returns a substring surrounding the first occurrence of phrase.
// The radius controls how many characters (runes) are kept on each side of the phrase,
// and omission specifies the string to append when content is truncated.
// If phrase is not found, the original string is returned unchanged.
func Excerpt(str, phrase string, options ...ExcerptOption) string {
	if phrase == "" {
		return str
	}

	opts := ExcerptOption{Radius: 100, Omission: "..."}
	if len(options) > 0 {
		if options[0].Radius != 0 {
			opts.Radius = options[0].Radius
		}
		if options[0].Omission != "" {
			opts.Omission = options[0].Omission
		}
	}

	radius := opts.Radius
	if radius < 0 {
		radius = 0
	}

	index := strings.Index(str, phrase)
	if index == -1 {
		return str
	}

	runes := []rune(str)
	phraseRunes := []rune(phrase)

	startRune := len([]rune(str[:index]))
	endRune := startRune + len(phraseRunes)

	prefixStart := startRune - radius
	if prefixStart < 0 {
		prefixStart = 0
	}

	suffixEnd := endRune + radius
	if suffixEnd > len(runes) {
		suffixEnd = len(runes)
	}

	prefix := string(runes[prefixStart:startRune])
	suffix := string(runes[endRune:suffixEnd])

	if prefixStart > 0 {
		prefix = strings.TrimLeft(prefix, " ")
		prefix = opts.Omission + prefix
	}

	if suffixEnd < len(runes) {
		suffix = strings.TrimRight(suffix, " ") + opts.Omission
	}

	return prefix + string(runes[startRune:endRune]) + suffix
}
