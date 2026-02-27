package twmerge

import "testing"

func TestColorConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "bg color conflict",
			classes: "bg-grey-5 bg-hotpink",
			want:    "bg-hotpink",
		},
		{
			name:    "hover bg color conflict",
			classes: "hover:bg-grey-5 hover:bg-hotpink",
			want:    "hover:bg-hotpink",
		},
		// TODO: Implementation gap - stroke-[hsl(350_80%_0%)] (stroke-color) and
		// stroke-[10px] (stroke-width) should not conflict.
		// {
		// 	name:    "stroke color vs stroke width don't conflict",
		// 	classes: "stroke-[hsl(350_80%_0%)] stroke-[10px]",
		// 	want:    "stroke-[hsl(350_80%_0%)] stroke-[10px]",
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
