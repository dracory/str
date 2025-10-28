package str

import "testing"

func TestReplace(t *testing.T) {
	tests := []struct {
		name          string
		in            string
		search        string
		replace       string
		caseSensitive []bool
		want          string
	}{
		{
			name:    "empty search returns original",
			in:      "hello",
			search:  "",
			replace: "world",
			want:    "hello",
		},
		{
			name:    "case sensitive replace",
			in:      "Hello hello",
			search:  "Hello",
			replace: "Hi",
			want:    "Hi hello",
		},
		{
			name:          "case insensitive replace",
			in:            "Hello hello",
			search:        "hello",
			replace:       "hi",
			caseSensitive: []bool{false},
			want:          "hi hi",
		},
		{
			name:    "no match returns original",
			in:      "abc",
			search:  "x",
			replace: "y",
			want:    "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.in, tt.search, tt.replace, tt.caseSensitive...); got != tt.want {
				t.Errorf("Replace() = %q, want %q", got, tt.want)
			}
		})
	}
}
