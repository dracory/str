package str

import "testing"

func TestWhenNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "non-empty string uses truthy",
			in:       "value",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "value!",
		},
		{
			name:     "empty string uses fallback",
			in:       "",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "empty string without fallback returns original",
			in:       "",
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenNotEmpty(tt.in, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenNotEmpty() = %q, want %q", got, tt.want)
			}
		})
	}
}
