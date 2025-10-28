package str

import "testing"

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		values []string
		want   bool
	}{
		{
			name:   "no values returns false",
			in:     "example",
			values: nil,
			want:   false,
		},
		{
			name:   "single match",
			in:     "example",
			values: []string{"ex"},
			want:   true,
		},
		{
			name:   "multiple values with match",
			in:     "example",
			values: []string{"no", "ex"},
			want:   true,
		},
		{
			name:   "multiple values without match",
			in:     "example",
			values: []string{"no", "pe"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartsWith(tt.in, tt.values...); got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
