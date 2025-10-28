package str

import "testing"

func TestContains(t *testing.T) {
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
			name:   "multiple values match first",
			input:  "golang",
			values: []string{"go", "lang"},
			want:   true,
		},
		{
			name:   "multiple values match later",
			input:  "golang",
			values: []string{"rust", "lang"},
			want:   true,
		},
		{
			name:   "no match",
			input:  "golang",
			values: []string{"python", "java"},
			want:   false,
		},
		{
			name:   "empty values ignored",
			input:  "golang",
			values: []string{"", "go"},
			want:   true,
		},
		{
			name:   "only empty values",
			input:  "golang",
			values: []string{"", " "},
			want:   false,
		},
		{
			name:   "empty input",
			input:  "",
			values: []string{"go"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.input, tt.values...); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
