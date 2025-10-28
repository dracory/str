package str

import "strings"

// PadLeft pads the string on the left to reach the specified length.
// When the pad argument is omitted or empty, spaces are used.
func PadLeft(str string, length int, pad ...string) string {
	defaultPad := " "
	if len(pad) > 0 && pad[0] != "" {
		defaultPad = pad[0]
	}

	short := Maximum(0, length-Length(str))
	if short == 0 {
		return str
	}

	leftPad := Substr(strings.Repeat(defaultPad, short), 0, short)
	return leftPad + str
}
