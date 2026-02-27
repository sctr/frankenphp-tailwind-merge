package twmerge

import "testing"

func TestArbitraryPropertyConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "handles arbitrary property conflicts correctly",
			classes: "[paint-order:markers] [paint-order:normal]",
			want:    "[paint-order:normal]",
		},
		{
			name:    "handles arbitrary property conflicts with multiple properties",
			classes: "[paint-order:markers] [--my-var:2rem] [paint-order:normal] [--my-var:4px]",
			want:    "[paint-order:normal] [--my-var:4px]",
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

func TestArbitraryPropertyConflictsWithModifiers(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "different modifiers don't conflict",
			classes: "[paint-order:markers] hover:[paint-order:normal]",
			want:    "[paint-order:markers] hover:[paint-order:normal]",
		},
		{
			name:    "same modifiers conflict",
			classes: "hover:[paint-order:markers] hover:[paint-order:normal]",
			want:    "hover:[paint-order:normal]",
		},
		{
			name:    "modifiers sorted correctly for conflict",
			classes: "hover:focus:[paint-order:markers] focus:hover:[paint-order:normal]",
			want:    "focus:hover:[paint-order:normal]",
		},
		{
			name:    "mixed properties with lg modifier",
			classes: "[paint-order:markers] [paint-order:normal] [--my-var:2rem] lg:[--my-var:4px]",
			want:    "[paint-order:normal] [--my-var:2rem] lg:[--my-var:4px]",
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

func TestArbitraryPropertyComplexConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "handles complex arbitrary property syntax",
			classes: "[-unknown-prop:::123:::] [-unknown-prop:url(https://hi.com)]",
			want:    "[-unknown-prop:url(https://hi.com)]",
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

func TestArbitraryPropertyImportantModifier(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "important doesn't conflict with non-important",
			classes: "![some:prop] [some:other]",
			want:    "![some:prop] [some:other]",
		},
		{
			name:    "important overrides previous important same property",
			classes: "![some:prop] [some:other] [some:one] ![some:another]",
			want:    "[some:one] ![some:another]",
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
