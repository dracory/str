package str

import "testing"

func TestPipe(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		callback func(string) string
		want     string
	}{
		{
			name:     "nil callback returns original",
			in:       "value",
			callback: nil,
			want:     "value",
		},
		{
			name: "uppercase transformation",
			in:   "go",
			callback: func(s string) string {
				return Upper(s)
			},
			want: "GO",
		},
		{
			name: "append suffix",
			in:   "str",
			callback: func(s string) string {
				return s + "ing"
			},
			want: "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pipe(tt.in, tt.callback); got != tt.want {
				t.Errorf("Pipe() = %q, want %q", got, tt.want)
			}
		})
	}
}
