package str

import "testing"

func TestWhenContainsAll(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		values   []string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "all values present uses truthy",
			in:       "hello brave new world",
			values:   []string{"hello", "world"},
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "hello brave new world!",
		},
		{
			name:     "missing value uses fallback",
			in:       "hello brave new world",
			values:   []string{"hello", "mars"},
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "missing value without fallback returns original",
			in:       "hello brave new world",
			values:   []string{"hello", "mars"},
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "hello brave new world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenContainsAll(tt.in, tt.values, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenContainsAll() = %q, want %q", got, tt.want)
			}
		})
	}
}
