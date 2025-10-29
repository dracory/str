package str

import "testing"

func TestTest(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		pattern string
		want    bool
	}{
		{
			name:    "matches pattern",
			in:      "abc123",
			pattern: "^\\w+$",
			want:    true,
		},
		{
			name:    "does not match pattern",
			in:      "abc-123",
			pattern: "^\\d+$",
			want:    false,
		},
		{
			name:    "empty pattern",
			in:      "anything",
			pattern: "",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Test(tt.in, tt.pattern); got != tt.want {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}
}
