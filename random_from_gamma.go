package str

import (
	"crypto/rand"
	"math/big"
)

// RandomFromGamma generates random string of specified length with the characters specified in the gamma string
func RandomFromGamma(length int, gamma string) string {
	if length <= 0 {
		return ""
	}

	if gamma == "" {
		return ""
	}

	inRune := []rune(gamma)
	out := make([]rune, length) // Pre-allocate for efficiency

	max := big.NewInt(int64(len(inRune)))
	for i := range out {
		randomIndex, err := rand.Int(rand.Reader, max)
		if err != nil {
			return ""
		}
		out[i] = inRune[randomIndex.Int64()]
	}

	return string(out)
}
