package str

import "testing"

func TestSnake(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		delimiter []string
		want      string
	}{
		{
			name: "default delimiter",
			in:   "HelloWorld",
			want: "hello_world",
		},
		{
			name:      "custom delimiter",
			in:        "HelloWorld",
			delimiter: []string{"-"},
			want:      "hello-world",
		},
		{
			name:      "empty delimiter treated as default",
			in:        "HelloWorld",
			delimiter: []string{""},
			want:      "hello_world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snake(tt.in, tt.delimiter...); got != tt.want {
				t.Errorf("Snake() = %q, want %q", got, tt.want)
			}
		})
	}
}
