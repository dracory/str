package str

import "testing"

func TestTitle(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"basic words", "hello world", "Hello World"},
		{"mixed case", "hELLo wORLD", "Hello World"},
		{"unicode", "mañana", "Mañana"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Title(tt.in); got != tt.want {
				t.Errorf("Title() = %q, want %q", got, tt.want)
			}
		})
	}
}
