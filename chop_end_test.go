package str

import "testing"

func TestChopEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		needle   string
		more     []string
		want     string
	}{
		{
			name:   "single match",
			input:  "foobar",
			needle: "bar",
			want:   "foo",
		},
		{
			name:   "first candidate matches",
			input:  "filename.txt",
			needle: ".md",
			more:   []string{".txt"},
			want:   "filename",
		},
		{
			name:   "second candidate matches",
			input:  "filename.txt",
			needle: ".log",
			more:   []string{".md", ".txt"},
			want:   "filename",
		},
		{
			name:   "no match returns original",
			input:  "filename.txt",
			needle: ".log",
			more:   []string{".md"},
			want:   "filename.txt",
		},
		{
			name:   "empty needle ignored",
			input:  "data.json",
			needle: "",
			more:   []string{".json"},
			want:   "data",
		},
		{
			name:   "empty string input",
			input:  "",
			needle: "a",
			want:   "",
		},
		{
			name:   "suffix appears multiple times",
			input:  "hellooo",
			needle: "o",
			want:   "helloo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChopEnd(tt.input, tt.needle, tt.more...)
			if got != tt.want {
				t.Errorf("ChopEnd() = %q, want %q", got, tt.want)
			}
		})
	}
}
