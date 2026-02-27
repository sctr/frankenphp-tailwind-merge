package twmerge

import "testing"

func TestPseudoVariantConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "empty pseudo variant conflict",
			classes: "empty:p-2 empty:p-3",
			want:    "empty:p-3",
		},
		{
			name:    "hover with empty pseudo variant conflict",
			classes: "hover:empty:p-2 hover:empty:p-3",
			want:    "hover:empty:p-3",
		},
		{
			name:    "read-only pseudo variant conflict",
			classes: "read-only:p-2 read-only:p-3",
			want:    "read-only:p-3",
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

func TestPseudoVariantGroupConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "group-empty conflict",
			classes: "group-empty:p-2 group-empty:p-3",
			want:    "group-empty:p-3",
		},
		{
			name:    "peer-empty conflict",
			classes: "peer-empty:p-2 peer-empty:p-3",
			want:    "peer-empty:p-3",
		},
		{
			name:    "group-empty and peer-empty don't conflict",
			classes: "group-empty:p-2 peer-empty:p-3",
			want:    "group-empty:p-2 peer-empty:p-3",
		},
		{
			name:    "hover with group-empty conflict",
			classes: "hover:group-empty:p-2 hover:group-empty:p-3",
			want:    "hover:group-empty:p-3",
		},
		{
			name:    "group-read-only conflict",
			classes: "group-read-only:p-2 group-read-only:p-3",
			want:    "group-read-only:p-3",
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
