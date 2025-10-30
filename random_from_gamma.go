package str

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// RandomFromGamma generates a random string of specified length using the characters provided in gamma.
// It returns an error when length is not positive, gamma is empty, or when secure randomness cannot be generated.
func RandomFromGamma(length int, gamma string) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than zero")
	}

	if gamma == "" {
		return "", fmt.Errorf("gamma must not be empty")
	}

	inRune := []rune(gamma)
	if len(inRune) == 0 {
		return "", fmt.Errorf("gamma must contain at least one valid rune")
	}

	out := make([]rune, length) // Pre-allocate for efficiency
	max := big.NewInt(int64(len(inRune)))

	for i := range out {
		randomIndex, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("crypto/rand failure: %w", err)
		}
		out[i] = inRune[randomIndex.Int64()]
	}

	return string(out), nil
}
