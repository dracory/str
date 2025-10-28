package str

import "testing"

func TestReplaceStart(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		search  string
		replace string
		want    string
	}{
		{
			name:   "empty search returns original",
			in:     "foobar",
			search: "",
			want:   "foobar",
		},
		{
			name:    "replace prefix",
			in:      "foobar",
			search:  "foo",
			replace: "baz",
			want:    "bazbar",
		},
		{
			name:    "no prefix match",
			in:      "foobar",
			search:  "bar",
			replace: "baz",
			want:    "foobar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceStart(tt.in, tt.search, tt.replace); got != tt.want {
				t.Errorf("ReplaceStart() = %q, want %q", got, tt.want)
			}
		})
	}
}
