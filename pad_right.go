package str

import "strings"

// PadRight pads the string on the right to reach the specified length.
// When the pad argument is omitted or empty, spaces are used.
func PadRight(str string, length int, pad ...string) string {
	defaultPad := " "
	if len(pad) > 0 && pad[0] != "" {
		defaultPad = pad[0]
	}

	short := Maximum(0, length-Length(str))
	if short == 0 {
		return str
	}

	rightPad := Substr(strings.Repeat(defaultPad, short), 0, short)
	return str + rightPad
}
