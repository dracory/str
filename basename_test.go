package str

import "testing"

func TestBasename(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		suffix []string
		want   string
	}{
		{
			name: "simple filename",
			path: "/path/to/file.txt",
			want: "file.txt",
		},
		{
			name:   "filename with suffix removal",
			path:   "/path/to/file.txt",
			suffix: []string{".txt"},
			want:   "file",
		},
		{
			name: "filename only",
			path: "file.txt",
			want: "file.txt",
		},
		{
			name:   "filename only with suffix removal",
			path:   "file.txt",
			suffix: []string{".txt"},
			want:   "file",
		},
		{
			name: "path with trailing slash",
			path: "/path/to/dir/",
			want: "dir",
		},
		{
			name: "empty path",
			path: "",
			want: ".",
		},
		{
			name:   "suffix not matching",
			path:   "file.txt",
			suffix: []string{".md"},
			want:   "file.txt",
		},
		{
			name:   "empty suffix",
			path:   "file.txt",
			suffix: []string{""},
			want:   "file.txt",
		},
		{
			name: "windows path",
			path: "C:\\path\\to\\file.txt",
			want: "file.txt",
		},
		{
			name:   "windows path with suffix",
			path:   "C:\\path\\to\\file.txt",
			suffix: []string{".txt"},
			want:   "file",
		},
		{
			name: "complex filename",
			path: "/var/www/html/index.html",
			want: "index.html",
		},
		{
			name:   "complex filename with suffix",
			path:   "/var/www/html/index.html",
			suffix: []string{".html"},
			want:   "index",
		},
		{
			name: "filename with multiple dots",
			path: "/path/to/file.tar.gz",
			want: "file.tar.gz",
		},
		{
			name:   "filename with multiple dots and suffix",
			path:   "/path/to/file.tar.gz",
			suffix: []string{".gz"},
			want:   "file.tar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Basename(tt.path, tt.suffix...)
			if got != tt.want {
				t.Errorf("Basename() = %v, want %v", got, tt.want)
			}
		})
	}
}
