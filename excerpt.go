package str

import (
	"strings"
	"unicode/utf8"
)

// ExcerptOption configures the behaviour of Excerpt.
type ExcerptOption struct {
	Radius   int
	Omission string
}

// Excerpt returns a substring surrounding the first occurrence of phrase.
// The radius controls how many runes are kept on each side of the phrase,
// and omission specifies the string to append when content is truncated.
// If phrase is not found, the original string is returned unchanged.
func Excerpt(str, phrase string, options ...ExcerptOption) string {
	if phrase == "" {
		return str
	}

	opts := ExcerptOption{Radius: 100, Omission: "..."}
	if len(options) > 0 {
		opts.Radius = options[0].Radius
		if options[0].Omission != "" {
			opts.Omission = options[0].Omission
		}
	}

	radius := opts.Radius
	if radius < 0 {
		radius = 0
	}

	idx := strings.Index(str, phrase)
	if idx == -1 {
		return str
	}

	before := str[:idx]
	after := str[idx+len(phrase):]

	startLen := utf8.RuneCountInString(before)
	startSlice := Substr(before, Maximum(startLen-radius, 0), radius)
	startBuilder := Of(startSlice)
	startBuilder = startBuilder.Unless(func(sb *StringBuilder) bool {
		return sb.Exactly(before)
	}, func(sb *StringBuilder) *StringBuilder {
		return sb.Prepend(opts.Omission)
	})

	endSlice := Substr(after, 0, radius)
	endBuilder := Of(endSlice)
	endBuilder = endBuilder.Unless(func(sb *StringBuilder) bool {
		return sb.Exactly(after)
	}, func(sb *StringBuilder) *StringBuilder {
		return sb.Append(opts.Omission)
	})

	return startBuilder.Append(phrase, endBuilder.String()).String()
}
