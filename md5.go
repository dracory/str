package str

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 converts a string to an MD5 hash.
//
// WARNING: MD5 is cryptographically broken and must not be used for security-
// sensitive purposes such as password hashing or integrity checks. This helper
// is retained for backward compatibility only; prefer SHA-256 via SHA256 (or stronger) for
// general hashing needs and BCrypt for passwords.
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
