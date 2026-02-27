package twmerge

import "testing"

func TestWonkyInputs(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "leading space",
			classes: " block",
			want:    "block",
		},
		{
			name:    "trailing space",
			classes: "block ",
			want:    "block",
		},
		{
			name:    "leading and trailing space",
			classes: " block ",
			want:    "block",
		},
		{
			name:    "multiple spaces between classes",
			classes: "  block  px-2     py-4  ",
			want:    "block px-2 py-4",
		},
		{
			name:    "newline between classes",
			classes: "block\npx-2",
			want:    "block px-2",
		},
		{
			name:    "newlines wrapping classes",
			classes: "\nblock\npx-2\n",
			want:    "block px-2",
		},
		{
			name:    "mixed whitespace types",
			classes: "  block\n        \n        px-2   \n          py-4  ",
			want:    "block px-2 py-4",
		},
		{
			name:    "carriage return and newlines",
			classes: "\r  block\n\r        \n        px-2   \n          py-4  ",
			want:    "block px-2 py-4",
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

func TestWonkyInputsMultipleArgs(t *testing.T) {
	tests := []struct {
		name    string
		classes []string
		want    string
	}{
		{
			name:    "multiple args with whitespace",
			classes: []string{"  block  px-2", " ", "     py-4  "},
			want:    "block px-2 py-4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwMerge(tt.classes...); got != tt.want {
				t.Errorf("TwMerge(%v) = %q, want %q", tt.classes, got, tt.want)
			}
		})
	}
}
