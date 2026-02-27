package twmerge

import "testing"

func TestNonTailwindClasses(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "non-tailwind class preserved",
			classes: "non-tailwind-class inline block",
			want:    "non-tailwind-class block",
		},
		{
			name:    "inline-like non-tailwind class preserved",
			classes: "inline block inline-1",
			want:    "block inline-1",
		},
		{
			name:    "prefix non-tailwind class preserved",
			classes: "inline block i-inline",
			want:    "block i-inline",
		},
		{
			name:    "non-tailwind with modifier preserved",
			classes: "focus:inline focus:block focus:inline-1",
			want:    "focus:block focus:inline-1",
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
