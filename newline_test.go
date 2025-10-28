package str

import "testing"

func TestNewLine(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		count []int
		want  string
	}{
		{
			name: "default single newline",
			in:   "line",
			want: "line\n",
		},
		{
			name:  "multiple newlines",
			in:    "line",
			count: []int{3},
			want:  "line\n\n\n",
		},
		{
			name:  "zero count adds nothing",
			in:    "line",
			count: []int{0},
			want:  "line",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLine(tt.in, tt.count...); got != tt.want {
				t.Errorf("NewLine() = %q, want %q", got, tt.want)
			}
		})
	}
}
