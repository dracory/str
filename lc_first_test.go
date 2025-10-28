package str

import "testing"

func TestLcFirst(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty string", "", ""},
		{"single rune upper", "Hello", "hello"},
		{"already lower", "world", "world"},
		{"unicode rune", "Ğo", "ğo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LcFirst(tt.in); got != tt.want {
				t.Errorf("LcFirst(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
