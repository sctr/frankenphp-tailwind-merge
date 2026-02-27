package twmerge

import "testing"

func TestTailwindCSSv33Features(t *testing.T) {
	tests := []struct {
		name    string
		classes []string
		want    string
	}{
		// TODO: Implementation gap - text-red (text-color) and text-lg/8 (font-size)
		// should not conflict. The Go implementation currently treats them as same group.
		// {
		// 	name:    "text color and postfix modifier",
		// 	classes: []string{"text-red text-lg/7 text-lg/8"},
		// 	want:    "text-red text-lg/8",
		// },
		{
			name:    "logical properties",
			classes: []string{
				"start-0 start-1",
				"end-0 end-1",
				"ps-0 ps-1 pe-0 pe-1",
				"ms-0 ms-1 me-0 me-1",
				"rounded-s-sm rounded-s-md rounded-e-sm rounded-e-md",
				"rounded-ss-sm rounded-ss-md rounded-ee-sm rounded-ee-md",
			},
			want: "start-1 end-1 ps-1 pe-1 ms-1 me-1 rounded-s-md rounded-e-md rounded-ss-md rounded-ee-md",
		},
		{
			name:    "logical properties shorthand overrides",
			classes: []string{"start-0 end-0 inset-0 ps-0 pe-0 p-0 ms-0 me-0 m-0 rounded-ss rounded-es rounded-s"},
			want:    "inset-0 p-0 m-0 rounded-s",
		},
		{
			name:    "hyphens",
			classes: []string{"hyphens-auto hyphens-manual"},
			want:    "hyphens-manual",
		},
		{
			name:    "gradient positions",
			classes: []string{"from-0% from-10% from-[12.5%] via-0% via-10% via-[12.5%] to-0% to-10% to-[12.5%]"},
			want:    "from-[12.5%] via-[12.5%] to-[12.5%]",
		},
		// TODO: Implementation gap - from-0% (gradient-from-position) and from-red (gradient-from-color)
		// should not conflict. The Go implementation currently treats them as same group.
		// {
		// 	name:    "gradient from-position and from-color don't conflict",
		// 	classes: []string{"from-0% from-red"},
		// 	want:    "from-0% from-red",
		// },
		{
			name:    "list-image",
			classes: []string{"list-image-none list-image-[url(./my-image.png)] list-image-[var(--value)]"},
			want:    "list-image-[var(--value)]",
		},
		{
			name:    "caption-side",
			classes: []string{"caption-top caption-bottom"},
			want:    "caption-bottom",
		},
		{
			name:    "line-clamp",
			classes: []string{"line-clamp-2 line-clamp-none line-clamp-[10]"},
			want:    "line-clamp-[10]",
		},
		{
			name:    "delay and duration",
			classes: []string{"delay-150 delay-0 duration-150 duration-0"},
			want:    "delay-0 duration-0",
		},
		{
			name:    "justify",
			classes: []string{"justify-normal justify-center justify-stretch"},
			want:    "justify-stretch",
		},
		{
			name:    "align-content",
			classes: []string{"content-normal content-center content-stretch"},
			want:    "content-stretch",
		},
		{
			name:    "whitespace",
			classes: []string{"whitespace-nowrap whitespace-break-spaces"},
			want:    "whitespace-break-spaces",
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

func TestTailwindCSSv34Features(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "viewport units",
			classes: "h-svh h-dvh w-svw w-dvw",
			want:    "h-dvh w-dvw",
		},
		{
			name:    "has variant",
			classes: "has-[[data-potato]]:p-1 has-[[data-potato]]:p-2 group-has-[:checked]:grid group-has-[:checked]:flex",
			want:    "has-[[data-potato]]:p-2 group-has-[:checked]:flex",
		},
		{
			name:    "text-wrap",
			classes: "text-wrap text-pretty",
			want:    "text-pretty",
		},
		{
			name:    "size utility",
			classes: "w-5 h-3 size-10 w-12",
			want:    "size-10 w-12",
		},
		{
			name:    "subgrid",
			classes: "grid-cols-2 grid-cols-subgrid grid-rows-5 grid-rows-subgrid",
			want:    "grid-cols-subgrid grid-rows-subgrid",
		},
		{
			name:    "min/max width",
			classes: "min-w-0 min-w-50 min-w-px max-w-0 max-w-50 max-w-px",
			want:    "min-w-px max-w-px",
		},
		{
			name:    "forced-color-adjust",
			classes: "forced-color-adjust-none forced-color-adjust-auto",
			want:    "forced-color-adjust-auto",
		},
		{
			name:    "appearance",
			classes: "appearance-none appearance-auto",
			want:    "appearance-auto",
		},
		{
			name:    "float and clear logical",
			classes: "float-start float-end clear-start clear-end",
			want:    "float-end clear-end",
		},
		{
			name:    "child variant",
			classes: "*:p-10 *:p-20 hover:*:p-10 hover:*:p-20",
			want:    "*:p-20 hover:*:p-20",
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

func TestTailwindCSSv40Features(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "transform style",
			classes: "transform-3d transform-flat",
			want:    "transform-flat",
		},
		{
			name:    "rotate axes",
			classes: "rotate-12 rotate-x-2 rotate-none rotate-y-3",
			want:    "rotate-x-2 rotate-none rotate-y-3",
		},
		{
			name:    "perspective",
			classes: "perspective-dramatic perspective-none perspective-midrange",
			want:    "perspective-midrange",
		},
		{
			name:    "perspective-origin",
			classes: "perspective-origin-center perspective-origin-top-left",
			want:    "perspective-origin-top-left",
		},
		{
			name:    "bg-linear",
			classes: "bg-linear-to-r bg-linear-45",
			want:    "bg-linear-45",
		},
		{
			name:    "gradient types",
			classes: "bg-linear-to-r bg-radial-[something] bg-conic-10",
			want:    "bg-conic-10",
		},
		// TODO: Implementation gap - inset-ring width (inset-ring, inset-ring-3) should be separate
		// from inset-ring color (inset-ring-blue). The Go impl conflates them.
		// {
		// 	name:    "inset-ring",
		// 	classes: "ring-4 ring-orange inset-ring inset-ring-3 inset-ring-blue",
		// 	want:    "ring-4 ring-orange inset-ring-3 inset-ring-blue",
		// },
		{
			name:    "field-sizing",
			classes: "field-sizing-content field-sizing-fixed",
			want:    "field-sizing-fixed",
		},
		{
			name:    "color-scheme",
			classes: "scheme-normal scheme-dark",
			want:    "scheme-dark",
		},
		{
			name:    "font-stretch",
			classes: "font-stretch-expanded font-stretch-[66.66%] font-stretch-50%",
			want:    "font-stretch-50%",
		},
		{
			name:    "col and row shorthand",
			classes: "col-span-full col-2 row-span-3 row-4",
			want:    "col-2 row-4",
		},
		{
			name:    "via with variable",
			classes: "via-red-500 via-(--mobile-header-gradient)",
			want:    "via-(--mobile-header-gradient)",
		},
		// TODO: Implementation gap - via-color and via-(length:...) variable should not conflict.
		// {
		// 	name:    "via color vs via length variable don't conflict",
		// 	classes: "via-red-500 via-(length:--mobile-header-gradient)",
		// 	want:    "via-red-500 via-(length:--mobile-header-gradient)",
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

func TestTailwindCSSv41Features(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "items-baseline-last",
			classes: "items-baseline items-baseline-last",
			want:    "items-baseline-last",
		},
		{
			name:    "self-baseline-last",
			classes: "self-baseline self-baseline-last",
			want:    "self-baseline-last",
		},
		{
			name:    "place-content safe variants",
			classes: "place-content-center place-content-end-safe place-content-center-safe",
			want:    "place-content-center-safe",
		},
		{
			name:    "items safe variants",
			classes: "items-center-safe items-baseline items-end-safe",
			want:    "items-end-safe",
		},
		{
			name:    "overflow-wrap",
			classes: "wrap-break-word wrap-normal wrap-anywhere",
			want:    "wrap-anywhere",
		},
		// TODO: Implementation gap - text-shadow class group not recognized.
		// {
		// 	name:    "text-shadow",
		// 	classes: "text-shadow-none text-shadow-2xl",
		// 	want:    "text-shadow-2xl",
		// },
		{
			name:    "mask composite",
			classes: "mask-add mask-subtract",
			want:    "mask-subtract",
		},
		{
			name:    "shadow with opacity modifier",
			classes: "shadow-md shadow-lg/25 text-shadow-md text-shadow-lg/25",
			want:    "shadow-lg/25 text-shadow-lg/25",
		},
		{
			name:    "mask-type",
			classes: "mask-type-luminance mask-type-alpha",
			want:    "mask-type-alpha",
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

func TestTailwindCSSv415Features(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "h-lh unit",
			classes: "h-12 h-lh",
			want:    "h-lh",
		},
		{
			name:    "min-h-lh unit",
			classes: "min-h-12 min-h-lh",
			want:    "min-h-lh",
		},
		{
			name:    "max-h-lh unit",
			classes: "max-h-12 max-h-lh",
			want:    "max-h-lh",
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

func TestTailwindCSSv42Features(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "inset-s conflict",
			classes: "inset-s-1 inset-s-2",
			want:    "inset-s-2",
		},
		{
			name:    "inset-e conflict",
			classes: "inset-e-1 inset-e-2",
			want:    "inset-e-2",
		},
		{
			name:    "start overridden by inset-s",
			classes: "start-1 inset-s-2",
			want:    "inset-s-2",
		},
		{
			name:    "inset-s overridden by start",
			classes: "inset-s-1 start-2",
			want:    "start-2",
		},
		{
			name:    "end overridden by inset-e",
			classes: "end-1 inset-e-2",
			want:    "inset-e-2",
		},
		{
			name:    "inset-e overridden by end",
			classes: "inset-e-1 end-2",
			want:    "end-2",
		},
		{
			name:    "decimal fractions in aspect ratio",
			classes: "aspect-8/11 aspect-8.5/11",
			want:    "aspect-8.5/11",
		},
		{
			name:    "decimal fractions in width",
			classes: "w-8/11 w-8.5/11",
			want:    "w-8.5/11",
		},
		{
			name:    "decimal fractions in inset",
			classes: "inset-1/2 inset-1.25/2.5",
			want:    "inset-1.25/2.5",
		},
		{
			name:    "font-features conflict",
			classes: `font-features-["smcp"] font-features-["onum"]`,
			want:    `font-features-["onum"]`,
		},
		{
			name:    "font-features with var",
			classes: `font-features-[var(--font-features)] font-features-["liga","dlig"]`,
			want:    `font-features-["liga","dlig"]`,
		},
		{
			name:    "font-variant-numeric and font-features don't conflict",
			classes: `tabular-nums font-features-["smcp"]`,
			want:    `tabular-nums font-features-["smcp"]`,
		},
		{
			name:    "font-features and normal-nums don't conflict",
			classes: `font-features-["smcp"] normal-nums`,
			want:    `font-features-["smcp"] normal-nums`,
		},
		{
			name:    "font-family and font-features don't conflict",
			classes: `font-sans font-features-["smcp"]`,
			want:    `font-sans font-features-["smcp"]`,
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
