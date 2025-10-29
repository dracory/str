package str

import "testing"

func TestWhenIsUuid(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "valid uuid uses truthy",
			in:       "123e4567-e89b-12d3-a456-426614174000",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "123e4567-e89b-12d3-a456-426614174000!",
		},
		{
			name:     "invalid uuid uses fallback",
			in:       "invalid",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "invalid uuid without fallback returns original",
			in:       "invalid",
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenIsUuid(tt.in, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenIsUuid() = %q, want %q", got, tt.want)
			}
		})
	}
}
