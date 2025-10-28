package str

import "regexp"

// Finish appends the suffix to str, ensuring it appears only once at the end.
func Finish(str, suffix string) string {
	if suffix == "" {
		return str
	}

	quoted := regexp.QuoteMeta(suffix)
	reg := regexp.MustCompile("(?:" + quoted + ")+$")

	return reg.ReplaceAllString(str, "") + suffix
}
