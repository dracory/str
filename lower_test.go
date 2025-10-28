package str

import "testing"

func TestLower(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"already lower", "lower", "lower"},
		{"upper to lower", "UPPER", "upper"},
		{"mixed case", "MiXeD", "mixed"},
		{"unicode", "Gö", "gö"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lower(tt.in); got != tt.want {
				t.Errorf("Lower(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
