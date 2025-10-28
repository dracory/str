package str

import "strings"

// Swap replaces occurrences of keys in the replacements map with their values.
// If the map is empty or contains an empty key, the original string is returned.
func Swap(str string, replacements map[string]string) string {
	if len(replacements) == 0 {
		return str
	}

	oldNewPairs := make([]string, 0, len(replacements)*2)
	for k, v := range replacements {
		if k == "" {
			return str
		}
		oldNewPairs = append(oldNewPairs, k, v)
	}

	replacer := strings.NewReplacer(oldNewPairs...)
	return replacer.Replace(str)
}
