package str

import "testing"

func TestWhenExactly(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		value    string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "exact match uses truthy",
			in:       "match",
			value:    "match",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "match!",
		},
		{
			name:     "no match uses fallback",
			in:       "match",
			value:    "other",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "no match without fallback returns original",
			in:       "match",
			value:    "other",
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "match",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenExactly(tt.in, tt.value, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenExactly() = %q, want %q", got, tt.want)
			}
		})
	}
}
