package str

import "testing"

func TestFinish(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		suffix string
		want   string
	}{
		{
			name:   "adds suffix when missing",
			str:    "path/to",
			suffix: "/",
			want:   "path/to/",
		},
		{
			name:   "does not duplicate existing suffix",
			str:    "path/to/",
			suffix: "/",
			want:   "path/to/",
		},
		{
			name:   "removes multiple trailing suffix occurrences",
			str:    "path///",
			suffix: "/",
			want:   "path/",
		},
		{
			name:   "empty suffix returns original",
			str:    "path",
			suffix: "",
			want:   "path",
		},
		{
			name:   "suffix with regex meta characters",
			str:    "file.txt.txt",
			suffix: ".txt",
			want:   "file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Finish(tt.str, tt.suffix); got != tt.want {
				t.Errorf("Finish() = %q, want %q", got, tt.want)
			}
		})
	}
}
