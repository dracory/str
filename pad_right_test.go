package str

import "testing"

func TestPadRight(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		length int
		pad    []string
		want   string
	}{
		{
			name:   "already long enough",
			in:     "golang",
			length: 3,
			want:   "golang",
		},
		{
			name:   "default space padding",
			in:     "go",
			length: 5,
			want:   "go   ",
		},
		{
			name:   "custom pad characters",
			in:     "go",
			length: 6,
			pad:    []string{"*"},
			want:   "go****",
		},
		{
			name:   "unicode aware length",
			in:     "été",
			length: 6,
			pad:    []string{"-"},
			want:   "été---",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRight(tt.in, tt.length, tt.pad...); got != tt.want {
				t.Errorf("PadRight() = %q, want %q", got, tt.want)
			}
		})
	}
}
