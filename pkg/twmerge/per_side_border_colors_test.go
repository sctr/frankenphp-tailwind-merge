package twmerge

import "testing"

func TestPerSideBorderColors(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "same side border color conflict",
			classes: "border-t-some-blue border-t-other-blue",
			want:    "border-t-other-blue",
		},
		{
			name:    "side border overridden by all-side border",
			classes: "border-t-some-blue border-some-blue",
			want:    "border-some-blue",
		},
		{
			name:    "all-side and logical side don't conflict",
			classes: "border-some-blue border-s-some-blue",
			want:    "border-some-blue border-s-some-blue",
		},
		{
			name:    "logical side border overridden by all-side border",
			classes: "border-e-some-blue border-some-blue",
			want:    "border-some-blue",
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
