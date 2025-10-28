package str

import "testing"

func TestRTrim(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		characters []string
		want       string
	}{
		{
			name: "default trims unicode space",
			in:   "hello\u00A0\t",
			want: "hello",
		},
		{
			name:       "custom characters",
			in:         "trim---",
			characters: []string{"-"},
			want:       "trim",
		},
		{
			name:       "empty charset behaves like default",
			in:         "spaced   ",
			characters: []string{""},
			want:       "spaced",
		},
		{
			name: "no trimming needed",
			in:   "value",
			want: "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RTrim(tt.in, tt.characters...); got != tt.want {
				t.Errorf("RTrim() = %q, want %q", got, tt.want)
			}
		})
	}
}
