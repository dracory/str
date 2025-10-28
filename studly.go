package str

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Studly converts a string to StudlyCase (also known as UpperCamelCase).
func Studly(in string) string {
	words := fieldsFunc(in, func(r rune) bool {
		return r == '_' || r == ' ' || r == '-' || r == ',' || r == '.'
	}, func(r rune) bool {
		return unicode.IsUpper(r)
	})

	title := cases.Title(language.Und)
	for i := range words {
		words[i] = title.String(words[i])
	}

	return strings.Join(words, "")
}
