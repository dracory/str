package str

import (
	"path/filepath"
	"testing"
)

func TestDirname(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		levels []int
		want   string
	}{
		{
			name: "single level",
			path: filepath.FromSlash("/path/to/file.txt"),
			want: filepath.FromSlash("/path/to"),
		},
		{
			name:   "two levels",
			path:   filepath.FromSlash("/path/to/file.txt"),
			levels: []int{2},
			want:   filepath.FromSlash("/path"),
		},
		{
			name:   "zero levels defaults to one",
			path:   filepath.FromSlash("/path/to/file.txt"),
			levels: []int{0},
			want:   filepath.FromSlash("/path/to"),
		},
		{
			name:   "negative levels ignored",
			path:   filepath.FromSlash("/path/to/file.txt"),
			levels: []int{-3},
			want:   filepath.FromSlash("/path/to"),
		},
		{
			name: "root stays root",
			path: filepath.Clean(filepath.FromSlash("/")),
			want: filepath.Clean(filepath.FromSlash("/")),
		},
		{
			name:   "windows path",
			path:   `C:\\path\\to\\file.txt`,
			levels: []int{1},
			want:   `C:\\path\\to`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dirname(tt.path, tt.levels...)
			if got != tt.want {
				t.Errorf("Dirname() = %q, want %q", got, tt.want)
			}
		})
	}
}
