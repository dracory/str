package str

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 returns the SHA-256 hash of the provided text encoded as a hex string.
func SHA256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
