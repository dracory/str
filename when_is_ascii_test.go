package str

import "testing"

func TestWhenIsAscii(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "ascii string uses truthy",
			in:       "hello",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "hello!",
		},
		{
			name:     "non-ascii string uses fallback",
			in:       "héllo",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "non-ascii without fallback returns original",
			in:       "héllo",
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "héllo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenIsAscii(tt.in, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenIsAscii() = %q, want %q", got, tt.want)
			}
		})
	}
}
