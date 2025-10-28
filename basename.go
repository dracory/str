package str

import (
	"path"
	"strings"
)

// Basename returns the basename of the file path string,
// and trims the suffix based on the parameter (optional).
// Example: Basename("/path/to/file.txt") returns "file.txt"
// Example: Basename("/path/to/file.txt", ".txt") returns "file"
func Basename(p string, suffix ...string) string {
	normalized := strings.ReplaceAll(p, "\\", "/")
	base := path.Base(normalized)
	if len(suffix) > 0 && suffix[0] != "" {
		base = strings.TrimSuffix(base, suffix[0])
	}
	return base
}
