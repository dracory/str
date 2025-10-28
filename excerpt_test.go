package str

import "testing"

func TestExcerpt(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		phrase  string
		options []ExcerptOption
		want    string
	}{
		{
			name:   "phrase not found returns original",
			str:    "The quick brown fox",
			phrase: "cat",
			want:   "The quick brown fox",
		},
		{
			name:   "no truncation with default radius",
			str:    "Hello brave new world",
			phrase: "brave",
			want:   "Hello brave new world",
		},
		{
			name:   "truncated both sides",
			str:    "The quick brown fox jumps over the lazy dog",
			phrase: "fox",
			options: []ExcerptOption{{Radius: 5}},
			want:   "...rown fox jump...",
		},
		{
			name:   "custom omission",
			str:    "Pack my box with five dozen liquor jugs",
			phrase: "five",
			options: []ExcerptOption{{Radius: 4, Omission: "~~~"}},
			want:   "~~~ith five doz~~~",
		},
		{
			name:   "zero radius keeps only phrase",
			str:    "Sphinx of black quartz, judge my vow",
			phrase: "black",
			options: []ExcerptOption{{Radius: 0}},
			want:   "...black...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Excerpt(tt.str, tt.phrase, tt.options...)
			if got != tt.want {
				t.Errorf("Excerpt() = %q, want %q", got, tt.want)
			}
		})
	}
}
