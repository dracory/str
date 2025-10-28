package str

import "testing"

func TestReplaceFirst(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		search  string
		replace string
		want    string
	}{
		{
			name:   "empty search returns original",
			in:     "hello world",
			search: "",
			want:   "hello world",
		},
		{
			name:    "replaces first occurrence",
			in:      "foo bar foo",
			search:  "foo",
			replace: "baz",
			want:    "baz bar foo",
		},
		{
			name:    "no match returns original",
			in:      "bar baz",
			search:  "foo",
			replace: "baz",
			want:    "bar baz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceFirst(tt.in, tt.search, tt.replace); got != tt.want {
				t.Errorf("ReplaceFirst() = %q, want %q", got, tt.want)
			}
		})
	}
}
