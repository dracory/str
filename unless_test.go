package str

import "testing"

func TestUnless(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		predicate func(string) bool
		fallback  func(string) string
		want      string
	}{
		{
			name:      "predicate true returns original",
			in:        "value",
			predicate: func(s string) bool { return s == "value" },
			fallback:  func(s string) string { return "fallback" },
			want:      "value",
		},
		{
			name:      "predicate false uses fallback",
			in:        "value",
			predicate: func(string) bool { return false },
			fallback:  func(string) string { return "fallback" },
			want:      "fallback",
		},
		{
			name:      "predicate false without fallback returns original",
			in:        "value",
			predicate: func(string) bool { return false },
			fallback:  nil,
			want:      "value",
		},
		{
			name:      "nil predicate returns original",
			in:        "value",
			predicate: nil,
			fallback:  func(string) string { return "fallback" },
			want:      "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unless(tt.in, tt.predicate, tt.fallback); got != tt.want {
				t.Errorf("Unless() = %q, want %q", got, tt.want)
			}
		})
	}
}
