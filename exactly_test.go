package str

import "testing"

func TestExactly(t *testing.T) {
	tests := []struct {
		name  string
		str   string
		value string
		want  bool
	}{
		{
			name:  "equal strings",
			str:   "hello",
			value: "hello",
			want:  true,
		},
		{
			name:  "different strings",
			str:   "hello",
			value: "world",
			want:  false,
		},
		{
			name:  "empty strings",
			str:   "",
			value: "",
			want:  true,
		},
		{
			name:  "case sensitive",
			str:   "Hello",
			value: "hello",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exactly(tt.str, tt.value); got != tt.want {
				t.Errorf("Exactly() = %v, want %v", got, tt.want)
			}
		})
	}
}
