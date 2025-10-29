package str

import "testing"

func TestWhenIsUlid(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		truthy   func(string) string
		fallback func(string) string
		want     string
	}{
		{
			name:     "valid ulid uses truthy",
			in:       "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "01ARZ3NDEKTSV4RRFFQ69G5FAV!",
		},
		{
			name:     "invalid ulid uses fallback",
			in:       "not-a-ulid",
			truthy:   func(s string) string { return s + "!" },
			fallback: func(string) string { return "fallback" },
			want:     "fallback",
		},
		{
			name:     "invalid ulid without fallback returns original",
			in:       "not-a-ulid",
			truthy:   func(s string) string { return s + "!" },
			fallback: nil,
			want:     "not-a-ulid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhenIsUlid(tt.in, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("WhenIsUlid() = %q, want %q", got, tt.want)
			}
		})
	}
}
