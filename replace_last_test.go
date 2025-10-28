package str

import "testing"

func TestReplaceLast(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		search  string
		replace string
		want    string
	}{
		{
			name:   "empty search returns original",
			in:     "foo",
			search: "",
			want:   "foo",
		},
		{
			name:    "replace last occurrence",
			in:      "foo bar foo",
			search:  "foo",
			replace: "baz",
			want:    "foo bar baz",
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
			if got := ReplaceLast(tt.in, tt.search, tt.replace); got != tt.want {
				t.Errorf("ReplaceLast() = %q, want %q", got, tt.want)
			}
		})
	}
}
