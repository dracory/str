package str

import "testing"

func TestSquish(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "leading and trailing whitespace",
			in:   "  hello  ",
			want: "hello",
		},
		{
			name: "multiple spaces collapsed",
			in:   "hello    world",
			want: "hello world",
		},
		{
			name: "tabs and newlines",
			in:   "hello\t\nworld",
			want: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Squish(tt.in); got != tt.want {
				t.Errorf("Squish() = %q, want %q", got, tt.want)
			}
		})
	}
}
