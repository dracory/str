package str

import "testing"

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		values []string
		want   bool
	}{
		{
			name:   "all present",
			input:  "hello world",
			values: []string{"hello", "world"},
			want:   true,
		},
		{
			name:   "missing one",
			input:  "hello world",
			values: []string{"hello", "golang"},
			want:   false,
		},
		{
			name:   "empty value treated as match",
			input:  "hello world",
			values: []string{"hello", ""},
			want:   true,
		},
		{
			name:   "empty input with empty value",
			input:  "",
			values: []string{""},
			want:   true,
		},
		{
			name:   "empty input with non-empty value",
			input:  "",
			values: []string{"hello"},
			want:   false,
		},
		{
			name:   "no values",
			input:  "hello",
			values: []string{},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.input, tt.values...); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
