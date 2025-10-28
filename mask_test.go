package str

import "testing"

func TestMask(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		char   string
		index  int
		length []int
		want   string
	}{
		{
			name:   "mask middle segment",
			in:     "secret123",
			char:   "*",
			index:  2,
			length: []int{3},
			want:   "se***t123",
		},
		{
			name:  "mask from end with negative index",
			in:    "secret123",
			char:  "#",
			index: -3,
			want:  "secret###",
		},
		{
			name:  "empty mask character returns original",
			in:    "secret",
			char:  "",
			index: 2,
			want:  "secret",
		},
		{
			name:  "segment outside bounds returns original",
			in:    "short",
			char:  "*",
			index: 10,
			want:  "short",
		},
		{
			name:   "unicode rune counting",
			in:     "héllo",
			char:   "*",
			index:  1,
			length: []int{2},
			want:   "h**lo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mask(tt.in, tt.char, tt.index, tt.length...); got != tt.want {
				t.Errorf("Mask() = %q, want %q", got, tt.want)
			}
		})
	}
}
