package str

import "testing"

func TestStudly(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"underscored words", "hello_world", "HelloWorld"},
		{"hyphenated words", "foo-bar", "FooBar"},
		{"spaced words", "foo bar", "FooBar"},
		{"mixed punctuation", "foo.bar,baz", "FooBarBaz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Studly(tt.in); got != tt.want {
				t.Errorf("Studly() = %q, want %q", got, tt.want)
			}
		})
	}
}
