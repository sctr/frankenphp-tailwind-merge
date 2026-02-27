package twmerge

import "testing"

func TestNegativeValueConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "negative margin conflicts",
			classes: "-m-2 -m-5",
			want:    "-m-5",
		},
		{
			name:    "negative top conflicts",
			classes: "-top-12 -top-2000",
			want:    "-top-2000",
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

func TestNegativeVsPositiveConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "negative margin then auto",
			classes: "-m-2 m-auto",
			want:    "m-auto",
		},
		{
			name:    "positive then negative top",
			classes: "top-12 -top-69",
			want:    "-top-69",
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

func TestNegativeCrossGroupConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "negative right overridden by inset-x",
			classes: "-right-1 inset-x-1",
			want:    "inset-x-1",
		},
		{
			name:    "negative with modifier stacks",
			classes: "hover:focus:-right-1 focus:hover:inset-x-1",
			want:    "focus:hover:inset-x-1",
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
