package str

import (
    "path"
    "strings"
)

// Dirname returns the directory portion of the provided path.
// By default it returns the parent directory; provide levels to traverse multiple parents.
func Dirname(p string, levels ...int) string {
    depth := 1
    if len(levels) > 0 {
        if levels[0] > 0 {
            depth = levels[0]
        }
    }

    normalized := strings.ReplaceAll(p, "\\", "/")
    dir := normalized
    for i := 0; i < depth; i++ {
        dir = path.Dir(dir)
    }

    if strings.Contains(p, "\\") {
        replacement := "\\"
        if strings.Contains(p, "\\\\") {
            replacement = "\\\\"
        }
        dir = strings.ReplaceAll(dir, "/", replacement)
    }

    return dir
}
