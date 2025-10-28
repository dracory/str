package str

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Title returns the input string converted to title case using the default
// language-independent rules.
func Title(in string) string {
	casesTitle := cases.Title(language.Und)
	return casesTitle.String(in)
}
