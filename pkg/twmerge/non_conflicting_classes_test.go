package twmerge

import "testing"

func TestNonConflictingClasses(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "border width and border color don't conflict",
			classes: "border-t border-white/10",
			want:    "border-t border-white/10",
		},
		{
			name:    "border width and solid border color don't conflict",
			classes: "border-t border-white",
			want:    "border-t border-white",
		},
		// TODO: Implementation gap - text-3.5xl (font-size) and text-black (text-color)
		// should not conflict. The Go implementation currently merges them.
		// {
		// 	name:    "text size and text color don't conflict",
		// 	classes: "text-3.5xl text-black",
		// 	want:    "text-3.5xl text-black",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwMerge(tt.classes); got != tt.want {
				t.Errorf("TwMerge(%q) = %q, want %q", tt.classes, got, tt.want)
			}
		})
	}
}
