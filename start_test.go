package str

import "testing"

func TestStart(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		prefix string
		want   string
	}{
		{
			name:   "empty prefix returns original",
			in:     "example",
			prefix: "",
			want:   "example",
		},
		{
			name:   "adds prefix when missing",
			in:     "path/to/file",
			prefix: "./",
			want:   "./path/to/file",
		},
		{
			name:   "does not duplicate existing prefix",
			in:     "./path/to/file",
			prefix: "./",
			want:   "./path/to/file",
		},
		{
			name:   "removes repeated prefix",
			in:     "http://http://example.com",
			prefix: "http://",
			want:   "http://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Start(tt.in, tt.prefix); got != tt.want {
				t.Errorf("Start() = %q, want %q", got, tt.want)
			}
		})
	}
}
