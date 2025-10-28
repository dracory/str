package str

import "strings"

// NewLine appends new line characters to the input string. When no count is
// provided, a single newline is appended. Otherwise, count[0] specifies how many
// newline characters to add.
func NewLine(in string, count ...int) string {
	if len(count) == 0 {
		return in + "\n"
	}

	return in + strings.Repeat("\n", count[0])
}
