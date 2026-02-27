package twmerge

import "testing"

func TestConflictsAcrossClassGroups(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "inset-1 then inset-x-1 both kept (shorthand first)",
			classes: "inset-1 inset-x-1",
			want:    "inset-1 inset-x-1",
		},
		{
			name:    "inset-x-1 then inset-1 collapses",
			classes: "inset-x-1 inset-1",
			want:    "inset-1",
		},
		{
			name:    "inset-x then left then inset collapses all",
			classes: "inset-x-1 left-1 inset-1",
			want:    "inset-1",
		},
		{
			name:    "inset-x then inset then left keeps inset + left",
			classes: "inset-x-1 inset-1 left-1",
			want:    "inset-1 left-1",
		},
		{
			name:    "inset-x then right then inset collapses",
			classes: "inset-x-1 right-1 inset-1",
			want:    "inset-1",
		},
		{
			name:    "inset-x then right then inset-x replaces",
			classes: "inset-x-1 right-1 inset-x-1",
			want:    "inset-x-1",
		},
		{
			name:    "inset-x inset-y don't conflict",
			classes: "inset-x-1 right-1 inset-y-1",
			want:    "inset-x-1 right-1 inset-y-1",
		},
		{
			name:    "right then inset-x then inset-y",
			classes: "right-1 inset-x-1 inset-y-1",
			want:    "inset-x-1 inset-y-1",
		},
		{
			name:    "hover modifier preserves across inset collapse",
			classes: "inset-x-1 hover:left-1 inset-1",
			want:    "hover:left-1 inset-1",
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

func TestRingAndShadowNoConflict(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "ring and shadow don't conflict",
			classes: "ring shadow",
			want:    "ring shadow",
		},
		{
			name:    "ring-2 and shadow-md don't conflict",
			classes: "ring-2 shadow-md",
			want:    "ring-2 shadow-md",
		},
		{
			name:    "shadow then ring don't conflict",
			classes: "shadow ring",
			want:    "shadow ring",
		},
		{
			name:    "shadow-md then ring-2 don't conflict",
			classes: "shadow-md ring-2",
			want:    "shadow-md ring-2",
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

func TestTouchClassConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "touch-pan-x then touch-pan-right",
			classes: "touch-pan-x touch-pan-right",
			want:    "touch-pan-right",
		},
		{
			name:    "touch-none then touch-pan-x",
			classes: "touch-none touch-pan-x",
			want:    "touch-pan-x",
		},
		{
			name:    "touch-pan-x then touch-none",
			classes: "touch-pan-x touch-none",
			want:    "touch-none",
		},
		{
			name:    "different touch actions don't conflict",
			classes: "touch-pan-x touch-pan-y touch-pinch-zoom",
			want:    "touch-pan-x touch-pan-y touch-pinch-zoom",
		},
		{
			name:    "touch-manipulation overridden by individual actions",
			classes: "touch-manipulation touch-pan-x touch-pan-y touch-pinch-zoom",
			want:    "touch-pan-x touch-pan-y touch-pinch-zoom",
		},
		{
			name:    "individual touch actions overridden by touch-auto",
			classes: "touch-pan-x touch-pan-y touch-pinch-zoom touch-auto",
			want:    "touch-auto",
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

func TestLineClampConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "line-clamp overrides overflow and display",
			classes: "overflow-auto inline line-clamp-1",
			want:    "line-clamp-1",
		},
		{
			name:    "line-clamp then overflow and display are kept",
			classes: "line-clamp-1 overflow-auto inline",
			want:    "line-clamp-1 overflow-auto inline",
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
