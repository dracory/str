package str

import "testing"

func TestWhenContains(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		value    string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "contains uses truthy",
			in:       "hello world",
			value:    "hello",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "hello world!",
		},
		{
			name:     "missing value uses fallback",
			in:       "hello world",
			value:    "goodbye",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "missing value without fallback returns original",
			in:       "hello world",
			value:    "goodbye",
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenContains(tt.in, tt.value, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenContains() = %q, want %q", got, tt.want)
			}
		})
	}
}
