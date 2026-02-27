package twmerge

import "testing"

func TestMergesClassesFromSameGroup(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "overflow-x conflict",
			classes: "overflow-x-auto overflow-x-hidden",
			want:    "overflow-x-hidden",
		},
		{
			name:    "basis conflict",
			classes: "basis-full basis-auto",
			want:    "basis-auto",
		},
		{
			name:    "width conflict",
			classes: "w-full w-fit",
			want:    "w-fit",
		},
		{
			name:    "three-way overflow-x conflict",
			classes: "overflow-x-auto overflow-x-hidden overflow-x-scroll",
			want:    "overflow-x-scroll",
		},
		{
			name:    "hover modifier preserves non-hover class",
			classes: "overflow-x-auto hover:overflow-x-hidden overflow-x-scroll",
			want:    "hover:overflow-x-hidden overflow-x-scroll",
		},
		{
			name:    "hover overrides previous hover",
			classes: "overflow-x-auto hover:overflow-x-hidden hover:overflow-x-auto overflow-x-scroll",
			want:    "hover:overflow-x-auto overflow-x-scroll",
		},
		{
			name:    "col-span conflict",
			classes: "col-span-1 col-span-full",
			want:    "col-span-full",
		},
		{
			name:    "multiple groups in one string",
			classes: "gap-2 gap-px basis-px basis-3",
			want:    "gap-px basis-3",
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

func TestFontVariantNumeric(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "different font variant numeric groups don't conflict",
			classes: "lining-nums tabular-nums diagonal-fractions",
			want:    "lining-nums tabular-nums diagonal-fractions",
		},
		{
			name:    "normal-nums overrides all font variant numeric",
			classes: "normal-nums tabular-nums diagonal-fractions",
			want:    "tabular-nums diagonal-fractions",
		},
		{
			name:    "normal-nums at end overrides all",
			classes: "tabular-nums diagonal-fractions normal-nums",
			want:    "normal-nums",
		},
		{
			name:    "same sub-group conflicts",
			classes: "tabular-nums proportional-nums",
			want:    "proportional-nums",
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
