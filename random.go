package str

import (
	"crypto/rand"
)

// Random generates an alphanumeric string of the specified length using crypto/rand
// with rejection sampling to keep the character distribution unbiased.
//
// Business Logic:
//  1. Directly using raw crypto/rand bytes (e.g. by base64-encoding and truncating)
//     skews character frequency, so we discard out-of-range samples before mapping.
//  2. Alphanumeric output (letters and digits only) keeps the resulting tokens URL,
//     filename, and copy/paste friendly.
func Random(length int) string {
	if length < 0 {
		panic("str.Random: negative length")
	}

	if length == 0 {
		return ""
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := len(charset)
	maxMultiple := (256 / charsetLen) * charsetLen

	result := make([]byte, length)
	for i := 0; i < length; {
		buf := make([]byte, length-i)
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}

		for _, b := range buf {
			if int(b) >= maxMultiple {
				continue
			}

			result[i] = charset[int(b)%charsetLen]
			i++

			if i == length {
				break
			}
		}
	}

	return string(result)
}
