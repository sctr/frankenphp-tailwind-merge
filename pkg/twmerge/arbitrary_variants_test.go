package twmerge

import "testing"

func TestBasicArbitraryVariants(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "basic element selector variant",
			classes: "[p]:underline [p]:line-through",
			want:    "[p]:line-through",
		},
		{
			name:    "child combinator variant",
			classes: "[&>*]:underline [&>*]:line-through",
			want:    "[&>*]:line-through",
		},
		{
			name:    "different variants don't conflict",
			classes: "[&>*]:underline [&>*]:line-through [&_div]:line-through",
			want:    "[&>*]:line-through [&_div]:line-through",
		},
		{
			name:    "supports variant",
			classes: "supports-[display:grid]:flex supports-[display:grid]:grid",
			want:    "supports-[display:grid]:grid",
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

func TestArbitraryVariantsWithModifiers(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "modifier chain with arbitrary variant",
			classes: "dark:lg:hover:[&>*]:underline dark:lg:hover:[&>*]:line-through",
			want:    "dark:lg:hover:[&>*]:line-through",
		},
		{
			name:    "reordered modifier chain with arbitrary variant",
			classes: "dark:lg:hover:[&>*]:underline dark:hover:lg:[&>*]:line-through",
			want:    "dark:hover:lg:[&>*]:line-through",
		},
		{
			name:    "position of arbitrary variant matters",
			classes: "hover:[&>*]:underline [&>*]:hover:line-through",
			want:    "hover:[&>*]:underline [&>*]:hover:line-through",
		},
		{
			name:    "complex mixed modifier positions",
			classes: "hover:dark:[&>*]:underline dark:hover:[&>*]:underline dark:[&>*]:hover:line-through",
			want:    "dark:hover:[&>*]:underline dark:[&>*]:hover:line-through",
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

func TestArbitraryVariantsComplexSyntax(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "nested @media in arbitrary variant",
			classes: "[@media_screen{@media(hover:hover)}]:underline [@media_screen{@media(hover:hover)}]:line-through",
			want:    "[@media_screen{@media(hover:hover)}]:line-through",
		},
		{
			name:    "modifier with nested @media arbitrary variant",
			classes: "hover:[@media_screen{@media(hover:hover)}]:underline hover:[@media_screen{@media(hover:hover)}]:line-through",
			want:    "hover:[@media_screen{@media(hover:hover)}]:line-through",
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

func TestArbitraryVariantsAttributeSelectors(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "data attribute selector",
			classes: "[&[data-open]]:underline [&[data-open]]:line-through",
			want:    "[&[data-open]]:line-through",
		},
		{
			name:    "multiple attribute selectors with pseudo class",
			classes: "[&[data-foo][data-bar]:not([data-baz])]:underline [&[data-foo][data-bar]:not([data-baz])]:line-through",
			want:    "[&[data-foo][data-bar]:not([data-baz])]:line-through",
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

func TestMultipleArbitraryVariants(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "same nested arbitrary variants conflict",
			classes: "[&>*]:[&_div]:underline [&>*]:[&_div]:line-through",
			want:    "[&>*]:[&_div]:line-through",
		},
		{
			name:    "different order nested arbitrary variants don't conflict",
			classes: "[&>*]:[&_div]:underline [&_div]:[&>*]:line-through",
			want:    "[&>*]:[&_div]:underline [&_div]:[&>*]:line-through",
		},
		{
			name:    "complex chain with multiple arbitrary variants",
			classes: "hover:dark:[&>*]:focus:disabled:[&_div]:underline dark:hover:[&>*]:disabled:focus:[&_div]:line-through",
			want:    "dark:hover:[&>*]:disabled:focus:[&_div]:line-through",
		},
		{
			name:    "different arbitrary variant positions prevent merge",
			classes: "hover:dark:[&>*]:focus:[&_div]:disabled:underline dark:hover:[&>*]:disabled:focus:[&_div]:line-through",
			want:    "hover:dark:[&>*]:focus:[&_div]:disabled:underline dark:hover:[&>*]:disabled:focus:[&_div]:line-through",
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

func TestArbitraryVariantsWithArbitraryProperties(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "arbitrary variant with arbitrary property",
			classes: "[&>*]:[color:red] [&>*]:[color:blue]",
			want:    "[&>*]:[color:blue]",
		},
		{
			name:    "complex selectors with arbitrary property",
			classes: "[&[data-foo][data-bar]:not([data-baz])]:nod:noa:[color:red] [&[data-foo][data-bar]:not([data-baz])]:noa:nod:[color:blue]",
			want:    "[&[data-foo][data-bar]:not([data-baz])]:noa:nod:[color:blue]",
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
