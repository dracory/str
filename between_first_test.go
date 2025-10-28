package str

import "testing"

func TestBetweenFirst(t *testing.T) {
	tests := []struct {
		name  string
		str   string
		start string
		end   string
		want  string
	}{
		{
			name:  "basic between first",
			str:   "Hello [World] and [Universe]",
			start: "[",
			end:   "]",
			want:  "World",
		},
		{
			name:  "multiple occurrences",
			str:   "a(b)c(d)e(f)",
			start: "(",
			end:   ")",
			want:  "b",
		},
		{
			name:  "nested brackets",
			str:   "outer [inner [nested] text] outer",
			start: "[",
			end:   "]",
			want:  "inner [nested",
		},
		{
			name:  "start not found",
			str:   "Hello World",
			start: "[",
			end:   "]",
			want:  "",
		},
		{
			name:  "end not found",
			str:   "Hello [World",
			start: "[",
			end:   "]",
			want:  "",
		},
		{
			name:  "empty start",
			str:   "Hello World",
			start: "",
			end:   "]",
			want:  "",
		},
		{
			name:  "empty end",
			str:   "Hello World",
			start: "[",
			end:   "",
			want:  "",
		},
		{
			name:  "empty string",
			str:   "",
			start: "[",
			end:   "]",
			want:  "",
		},
		{
			name:  "start and end same",
			str:   "Hello 'World' and 'Universe'",
			start: "'",
			end:   "'",
			want:  "World",
		},
		{
			name:  "multi-char delimiters",
			str:   "Hello {{World}} and {{Universe}}",
			start: "{{",
			end:   "}}",
			want:  "World",
		},
		{
			name:  "start at beginning",
			str:   "[Hello] World",
			start: "[",
			end:   "]",
			want:  "Hello",
		},
		{
			name:  "end at end",
			str:   "Hello [World]",
			start: "[",
			end:   "]",
			want:  "World",
		},
		{
			name:  "empty content between",
			str:   "Hello [] World",
			start: "[",
			end:   "]",
			want:  "",
		},
		{
			name:  "whitespace content",
			str:   "Hello [   ] World",
			start: "[",
			end:   "]",
			want:  "   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BetweenFirst(tt.str, tt.start, tt.end)
			if got != tt.want {
				t.Errorf("BetweenFirst() = %q, want %q", got, tt.want)
			}
		})
	}
}
