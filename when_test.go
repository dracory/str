package str

import "testing"

func TestWhen(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		condition bool
		truthy    func(string) string
		fallback  func(string) string
		want      string
	}{
		{
			name:      "condition true uses truthy",
			in:        "value",
			condition: true,
			truthy:    func(s string) string { return s + "!" },
			fallback:  func(string) string { return "fallback" },
			want:      "value!",
		},
		{
			name:      "condition true without truthy returns original",
			in:        "value",
			condition: true,
			truthy:    nil,
			fallback:  func(string) string { return "fallback" },
			want:      "value",
		},
		{
			name:      "condition false with fallback",
			in:        "value",
			condition: false,
			truthy:    func(string) string { return "ignored" },
			fallback:  func(string) string { return "fallback" },
			want:      "fallback",
		},
		{
			name:      "condition false without fallback returns original",
			in:        "value",
			condition: false,
			truthy:    func(string) string { return "ignored" },
			fallback:  nil,
			want:      "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := When(tt.in, tt.condition, tt.truthy, tt.fallback); got != tt.want {
				t.Errorf("When() = %q, want %q", got, tt.want)
			}
		})
	}
}
