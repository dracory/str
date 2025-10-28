package str

import "testing"

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		values []string
		want   bool
	}{
		{
			name:   "single match",
			input:  "hello world",
			values: []string{"world"},
			want:   true,
		},
		{
			name:   "multiple values match",
			input:  "filename.txt",
			values: []string{".log", ".txt"},
			want:   true,
		},
		{
			name:   "no match",
			input:  "filename.txt",
			values: []string{".md", ".log"},
			want:   false,
		},
		{
			name:   "empty values ignored",
			input:  "filename.txt",
			values: []string{"", ".txt"},
			want:   true,
		},
		{
			name:   "only empty values",
			input:  "filename.txt",
			values: []string{"", " "},
			want:   false,
		},
		{
			name:   "empty input",
			input:  "",
			values: []string{"suffix"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndsWith(tt.input, tt.values...); got != tt.want {
				t.Errorf("EndsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
