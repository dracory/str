package str

import "path/filepath"

// Dirname returns the directory portion of the provided path.
// By default it returns the parent directory; provide levels to traverse multiple parents.
func Dirname(path string, levels ...int) string {
	depth := 1
	if len(levels) > 0 {
		if levels[0] > 0 {
			depth = levels[0]
		}
	}

	dir := path
	for i := 0; i < depth; i++ {
		dir = filepath.Dir(dir)
	}

	return dir
}
