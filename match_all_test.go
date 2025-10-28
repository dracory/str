package str

import "testing"

func TestMatchAll(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		pattern string
		want    []string
	}{
		{
			name:    "empty pattern returns whole string",
			in:      "hello",
			pattern: "",
			want:    []string{"hello"},
		},
		{
			name:    "captures all digits",
			in:      "foo123bar456",
			pattern: `\d+`,
			want:    []string{"123", "456"},
		},
		{
			name:    "no matches returns empty slice",
			in:      "abcdef",
			pattern: `\d+`,
			want:    []string{},
		},
		{
			name:    "unicode words",
			in:      "Привет мир 123",
			pattern: `\p{L}+`,
			want:    []string{"Привет", "мир"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatchAll(tt.in, tt.pattern)
			if len(got) != len(tt.want) {
				t.Fatalf("MatchAll() length = %d, want %d, got %#v", len(got), len(tt.want), got)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("MatchAll()[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}
