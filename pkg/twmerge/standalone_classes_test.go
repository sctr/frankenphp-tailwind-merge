package twmerge

import "testing"

func TestStandaloneClassConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "inline block conflict",
			classes: "inline block",
			want:    "block",
		},
		{
			name:    "hover modifier conflict",
			classes: "hover:block hover:inline",
			want:    "hover:inline",
		},
		{
			name:    "duplicate class dedup",
			classes: "hover:block hover:block",
			want:    "hover:block",
		},
		{
			name:    "complex modifier chain with standalone classes",
			classes: "inline hover:inline focus:inline hover:block hover:focus:block",
			want:    "inline focus:inline hover:block hover:focus:block",
		},
		{
			name:    "text decoration conflict",
			classes: "underline line-through",
			want:    "line-through",
		},
		{
			name:    "text decoration with no-underline",
			classes: "line-through no-underline",
			want:    "no-underline",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwMerge(tt.classes); got != tt.want {
				t.Errorf("TwMerge(%q) = %q, want %q", tt.classes, got, tt.want)
			}
		})
	}
}
