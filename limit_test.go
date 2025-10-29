package str

import "testing"

func TestLimit(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		limit int
		end   []string
		want  string
	}{
		{
			name:  "no truncation",
			in:    "hello",
			limit: 10,
			want:  "hello",
		},
		{
			name:  "default suffix",
			in:    "hello world",
			limit: 5,
			want:  "hello...",
		},
		{
			name:  "custom suffix",
			in:    "hello world",
			limit: 5,
			end:   []string{"~"},
			want:  "hello~",
		},
		{
			name:  "limit zero",
			in:    "hello",
			limit: 0,
			want:  "...",
		},
		{
			name:  "unicode aware",
			in:    "héllö",
			limit: 3,
			want:  "hél...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Limit(tt.in, tt.limit, tt.end...); got != tt.want {
				t.Errorf("Limit() = %q, want %q", got, tt.want)
			}
		})
	}
}
