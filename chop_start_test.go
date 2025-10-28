package str

import "testing"

func TestChopStart(t *testing.T) {
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
			needle: "foo",
			want:   "bar",
		},
		{
			name:   "first candidate matches",
			input:  "log_filename",
			needle: "error_",
			more:   []string{"log_"},
			want:   "filename",
		},
		{
			name:   "second candidate matches",
			input:  "log_filename",
			needle: "debug_",
			more:   []string{"error_", "log_"},
			want:   "filename",
		},
		{
			name:   "no match returns original",
			input:  "filename",
			needle: "log_",
			more:   []string{"debug_"},
			want:   "filename",
		},
		{
			name:   "empty needle ignored",
			input:  "data_file",
			needle: "",
			more:   []string{"data_"},
			want:   "file",
		},
		{
			name:   "empty string input",
			input:  "",
			needle: "a",
			want:   "",
		},
		{
			name:   "prefix appears multiple times",
			input:  "prefixprefixvalue",
			needle: "prefix",
			want:   "prefixvalue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChopStart(tt.input, tt.needle, tt.more...)
			if got != tt.want {
				t.Errorf("ChopStart() = %q, want %q", got, tt.want)
			}
		})
	}
}
