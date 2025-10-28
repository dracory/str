package str

import "testing"

func TestReplaceMatches(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		pattern string
		replace string
		want    string
	}{
		{
			name:    "empty pattern returns original",
			in:      "abc123",
			pattern: "",
			replace: "-",
			want:    "abc123",
		},
		{
			name:    "replace digits",
			in:      "abc123def456",
			pattern: `\d+`,
			replace: "-",
			want:    "abc-def-",
		},
		{
			name:    "unicode letters",
			in:      "Привет123",
			pattern: `\p{L}+`,
			replace: "hi",
			want:    "hi123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceMatches(tt.in, tt.pattern, tt.replace); got != tt.want {
				t.Errorf("ReplaceMatches() = %q, want %q", got, tt.want)
			}
		})
	}
}
