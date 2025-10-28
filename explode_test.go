package str

import (
	"testing"
)

func TestExplode(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		delimiter string
		limit     []int
		want      []string
	}{
		{
			name:      "no limit provided",
			str:       "a,b,c",
			delimiter: ",",
			want:      []string{"a", "b", "c"},
		},
		{
			name:      "limit greater than parts",
			str:       "a,b",
			delimiter: ",",
			limit:     []int{5},
			want:      []string{"a", "b"},
		},
		{
			name:      "positive limit splits remainder",
			str:       "a,b,c,d",
			delimiter: ",",
			limit:     []int{3},
			want:      []string{"a", "b", "c,d"},
		},
		{
			name:      "negative limit removes tail",
			str:       "a,b,c,d",
			delimiter: ",",
			limit:     []int{-1},
			want:      []string{"a", "b", "c"},
		},
		{
			name:      "negative limit removing entire result",
			str:       "a,b",
			delimiter: ",",
			limit:     []int{-2},
			want:      []string{},
		},
		{
			name:      "limit zero treated as unlimited",
			str:       "a|b|c",
			delimiter: "|",
			limit:     []int{0},
			want:      []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Explode(tt.str, tt.delimiter, tt.limit...)
			if len(got) != len(tt.want) {
				t.Fatalf("Explode() length = %d, want %d, got: %#v", len(got), len(tt.want), got)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Explode()[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}
