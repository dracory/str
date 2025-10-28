package str

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		times int
		want  string
	}{
		{"zero times returns empty", "go", 0, ""},
		{"negative times returns empty", "go", -2, ""},
		{"repeat once", "go", 1, "go"},
		{"repeat multiple", "go", 3, "gogogo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.in, tt.times); got != tt.want {
				t.Errorf("Repeat() = %q, want %q", got, tt.want)
			}
		})
	}
}
