package str

import "testing"

func TestLength(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"empty", "", 0},
		{"ascii", "hello", 5},
		{"unicode", "héllo", 5},
		{"emoji", "👋🌍", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Length(tt.in); got != tt.want {
				t.Errorf("Length(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
