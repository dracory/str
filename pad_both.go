package str

import "strings"

// PadBoth pads the string on both sides to reach the desired length.
// Padding is evenly split between the left and right; when an odd number of
// characters is required, the extra character is added to the right side.
// The optional pad argument specifies the padding characters; it defaults to a
// single space when omitted or empty.
func PadBoth(str string, length int, pad ...string) string {
	defaultPad := " "
	if len(pad) > 0 && pad[0] != "" {
		defaultPad = pad[0]
	}

	short := Maximum(0, length-Length(str))
	if short == 0 {
		return str
	}

	left := short / 2
	right := short/2 + short%2

	leftPad := Substr(strings.Repeat(defaultPad, left), 0, left)
	rightPad := Substr(strings.Repeat(defaultPad, right), 0, right)

	return leftPad + str + rightPad
}
