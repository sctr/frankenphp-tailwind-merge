package twmerge

import "testing"

func TestImportantModifierV4(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "v4 suffix important conflict",
			classes: "font-medium! font-bold!",
			want:    "font-bold!",
		},
		{
			name:    "important and non-important both kept",
			classes: "font-medium! font-bold! font-thin",
			want:    "font-bold! font-thin",
		},
		{
			name:    "important inset cross-group conflict",
			classes: "right-2! -inset-x-px!",
			want:    "-inset-x-px!",
		},
		{
			name:    "important with focus modifier",
			classes: "focus:inline! focus:block!",
			want:    "focus:block!",
		},
		{
			name:    "important with CSS variable arbitrary property",
			classes: "[--my-var:20px]! [--my-var:30px]!",
			want:    "[--my-var:30px]!",
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

func TestImportantModifierV3Legacy(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "v3 prefix important conflict",
			classes: "!font-medium !font-bold",
			want:    "!font-bold",
		},
		{
			name:    "v3 important and non-important both kept",
			classes: "!font-medium !font-bold font-thin",
			want:    "!font-bold font-thin",
		},
		{
			name:    "v3 important inset cross-group conflict",
			classes: "!right-2 !-inset-x-px",
			want:    "!-inset-x-px",
		},
		{
			name:    "v3 important with focus modifier",
			classes: "focus:!inline focus:!block",
			want:    "focus:!block",
		},
		{
			name:    "v3 important with CSS variable arbitrary property",
			classes: "![--my-var:20px] ![--my-var:30px]",
			want:    "![--my-var:30px]",
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

func TestImportantModifierMixedV3V4(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "v4 suffix overridden by v3 prefix",
			classes: "font-medium! !font-bold",
			want:    "!font-bold",
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
