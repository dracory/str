package str

import "testing"

func TestPadBoth(t *testing.T) {
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
			name:   "even padding with default space",
			in:     "go",
			length: 10,
			want:   "    go    ",
		},
		{
			name:   "odd padding distributes extra to right",
			in:     "go",
			length: 5,
			want:   " go  ",
		},
		{
			name:   "custom pad characters",
			in:     "go",
			length: 6,
			pad:    []string{"*"},
			want:   "**go**",
		},
		{
			name:   "unicode aware length",
			in:     "été",
			length: 7,
			pad:    []string{"-"},
			want:   "--été--",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadBoth(tt.in, tt.length, tt.pad...); got != tt.want {
				t.Errorf("PadBoth() = %q, want %q", got, tt.want)
			}
		})
	}
}
