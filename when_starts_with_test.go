package str

import "testing"

func TestWhenStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		values   []string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "matches values uses truthy",
			in:       "prefix-value",
			values:   []string{"pre", "foo"},
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "prefix-value!",
		},
		{
			name:     "no match uses fallback",
			in:       "value",
			values:   []string{"pre"},
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "no match without fallback returns original",
			in:       "value",
			values:   []string{"pre"},
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenStartsWith(tt.in, tt.values, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenStartsWith() = %q, want %q", got, tt.want)
			}
		})
	}
}
