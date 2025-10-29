package str

import "testing"

func TestWhenEmpty(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "empty string uses truthy",
			in:       "",
			truthy:   func(s string) string { return "empty" },
			fallback: func(string) string { return "fallback" },
			want:     "empty",
		},
		{
			name:     "non-empty string uses fallback",
			in:       "value",
			truthy:   func(string) string { return "empty" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "non-empty without fallback returns original",
			in:       "value",
			truthy:   func(string) string { return "empty" },
			fallback: nil,
			want:     "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenEmpty(tt.in, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenEmpty() = %q, want %q", got, tt.want)
			}
		})
	}
}
