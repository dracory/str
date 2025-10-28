package str

import "testing"

func TestHeadline(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "already spaced becomes title case",
			in:   "hello world",
			want: "Hello World",
		},
		{
			name: "single word studly split",
			in:   "foo_bar",
			want: "Foo Bar",
		},
		{
			name: "handles hyphenated",
			in:   "foo-bar",
			want: "Foo Bar",
		},
		{
			name: "preserves acronyms",
			in:   "HTTPServer",
			want: "HTTP Server",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Headline(tt.in); got != tt.want {
				t.Errorf("Headline() = %q, want %q", got, tt.want)
			}
		})
	}
}
