package str

import "testing"

func TestRemove(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		values []string
		want   string
	}{
		{
			name: "no values returns original",
			in:   "hello",
			want: "hello",
		},
		{
			name:   "single removal",
			in:     "hello world",
			values: []string{" world"},
			want:   "hello",
		},
		{
			name:   "multiple removals",
			in:     "foo bar baz",
			values: []string{"foo ", " baz"},
			want:   "bar",
		},
		{
			name:   "overlapping values",
			in:     "banana",
			values: []string{"a", "n"},
			want:   "b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.in, tt.values...); got != tt.want {
				t.Errorf("Remove() = %q, want %q", got, tt.want)
			}
		})
	}
}
