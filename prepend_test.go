package str

import "testing"

func TestPrepend(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		values []string
		want   string
	}{
		{
			name: "no values returns original",
			in:   "world",
			want: "world",
		},
		{
			name:   "single value",
			in:     "world",
			values: []string{"hello "},
			want:   "hello world",
		},
		{
			name:   "multiple values",
			in:     "world",
			values: []string{"hello", " ", "beautiful "},
			want:   "hello beautiful world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prepend(tt.in, tt.values...); got != tt.want {
				t.Errorf("Prepend() = %q, want %q", got, tt.want)
			}
		})
	}
}
