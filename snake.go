package str

import "strings"

// Snake converts the string to snake case using the provided delimiter.
// The delimiter defaults to underscore when not specified or empty.
func Snake(str string, delimiter ...string) string {
	result := ToSnake(str)
	del := "_"
	if len(delimiter) > 0 && delimiter[0] != "" {
		del = delimiter[0]
	}

	if del == "_" {
		return result
	}

	return strings.ReplaceAll(result, "_", del)
}
