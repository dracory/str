package str

import "testing"

func TestSwap(t *testing.T) {
	tests := []struct {
		name         string
		in           string
		replacements map[string]string
		want         string
	}{
		{
			name:         "empty map returns original",
			in:           "hello",
			replacements: map[string]string{},
			want:         "hello",
		},
		{
			name: "single replacement",
			in:   "hello",
			replacements: map[string]string{
				"hello": "hi",
			},
			want: "hi",
		},
		{
			name: "multiple replacements",
			in:   "foo bar baz",
			replacements: map[string]string{
				"foo": "qux",
				"bar": "quux",
			},
			want: "qux quux baz",
		},
		{
			name: "empty key returns original",
			in:   "hello",
			replacements: map[string]string{
				"": "hi",
			},
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Swap(tt.in, tt.replacements); got != tt.want {
				t.Errorf("Swap() = %q, want %q", got, tt.want)
			}
		})
	}
}
