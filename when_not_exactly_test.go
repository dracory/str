package str

import "testing"

func TestWhenNotExactly(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		value    string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "different value uses truthy",
			in:       "first",
			value:    "second",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "first!",
		},
		{
			name:     "same value uses fallback",
			in:       "first",
			value:    "first",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "same value without fallback returns original",
			in:       "first",
			value:    "first",
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "first",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenNotExactly(tt.in, tt.value, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenNotExactly() = %q, want %q", got, tt.want)
			}
		})
	}
}
