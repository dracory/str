package str

import "testing"

func TestSplit(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		pattern string
		limit   []int
		want    []string
	}{
		{
			name:    "empty pattern returns whole string",
			in:      "a,b,c",
			pattern: "",
			want:    []string{"a,b,c"},
		},
		{
			name:    "split by comma",
			in:      "a,b,c",
			pattern: ",",
			want:    []string{"a", "b", "c"},
		},
		{
			name:    "respect limit",
			in:      "a|b|c|d",
			pattern: "\\|",
			limit:   []int{3},
			want:    []string{"a", "b", "c|d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.in, tt.pattern, tt.limit...)
			if len(got) != len(tt.want) {
				t.Fatalf("Split() length = %d, want %d, got %#v", len(got), len(tt.want), got)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("Split()[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}
