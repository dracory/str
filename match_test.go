package str

import "testing"

func TestMatch(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		pattern string
		want    string
	}{
		{
			name:    "empty pattern returns original",
			in:      "hello",
			pattern: "",
			want:    "hello",
		},
		{
			name:    "matches first occurrence",
			in:      "foo123bar456",
			pattern: `\d+`,
			want:    "123",
		},
		{
			name:    "no match returns empty",
			in:      "abcdef",
			pattern: `\d+`,
			want:    "",
		},
		{
			name:    "unicode match",
			in:      "Привет123",
			pattern: `\p{L}+`,
			want:    "Привет",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.in, tt.pattern); got != tt.want {
				t.Errorf("Match() = %q, want %q", got, tt.want)
			}
		})
	}
}
