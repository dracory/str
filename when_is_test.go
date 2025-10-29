package str

import "testing"

func TestWhenIs(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		patterns []string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "pattern matches uses truthy",
			in:       "example.txt",
			patterns: []string{"*.txt", "*.md"},
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "example.txt!",
		},
		{
			name:     "no pattern match uses fallback",
			in:       "example.txt",
			patterns: []string{"*.md"},
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "no pattern match without fallback returns original",
			in:       "example.txt",
			patterns: []string{"*.md"},
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "example.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenIs(tt.in, tt.patterns, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenIs() = %q, want %q", got, tt.want)
			}
		})
	}
}
