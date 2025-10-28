package str

import "testing"

func TestTrim(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		characters []string
		want       string
	}{
		{
			name: "default trim",
			in:   "  hello\n",
			want: "hello",
		},
		{
			name:       "custom characters",
			in:         "---hello---",
			characters: []string{"-"},
			want:       "hello",
		},
		{
			name:       "empty characters behaves like default",
			in:         "  hello  ",
			characters: []string{""},
			want:       "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.in, tt.characters...); got != tt.want {
				t.Errorf("Trim() = %q, want %q", got, tt.want)
			}
		})
	}
}
