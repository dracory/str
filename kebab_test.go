package str

import "testing"

func TestKebab(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "basic",
			in:   "HelloWorld",
			want: "hello-world",
		},
		{
			name: "with spaces",
			in:   "Hello World Example",
			want: "hello-world-example",
		},
		{
			name: "already kebab",
			in:   "already-kebab",
			want: "already-kebab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kebab(tt.in); got != tt.want {
				t.Errorf("Kebab() = %q, want %q", got, tt.want)
			}
		})
	}
}
