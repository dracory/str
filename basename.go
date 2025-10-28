package str

import (
	"path/filepath"
	"strings"
)

// Basename returns the basename of the file path string,
// and trims the suffix based on the parameter (optional).
// Example: Basename("/path/to/file.txt") returns "file.txt"
// Example: Basename("/path/to/file.txt", ".txt") returns "file"
func Basename(path string, suffix ...string) string {
	base := filepath.Base(path)
	if len(suffix) > 0 && suffix[0] != "" {
		base = strings.TrimSuffix(base, suffix[0])
	}
	return base
}
