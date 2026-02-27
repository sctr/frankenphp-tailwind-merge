package twmerge

import "testing"

func TestArbitraryValueConflicts(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "basic arbitrary value conflict",
			classes: "m-[2px] m-[10px]",
			want:    "m-[10px]",
		},
		{
			name:    "arbitrary value with various units",
			classes: "m-[2px] m-[11svmin] m-[12in] m-[13lvi] m-[14vb] m-[15vmax] m-[16mm] m-[17%] m-[18em] m-[19px] m-[10dvh]",
			want:    "m-[10dvh]",
		},
		{
			name:    "arbitrary value with container query units",
			classes: "h-[10px] h-[11cqw] h-[12cqh] h-[13cqi] h-[14cqb] h-[15cqmin] h-[16cqmax]",
			want:    "h-[16cqmax]",
		},
		{
			name:    "z-index with arbitrary value",
			classes: "z-20 z-[99]",
			want:    "z-[99]",
		},
		{
			name:    "cross-group arbitrary merge my to m",
			classes: "my-[2px] m-[10rem]",
			want:    "m-[10rem]",
		},
		{
			name:    "cursor with arbitrary value",
			classes: "cursor-pointer cursor-[grab]",
			want:    "cursor-[grab]",
		},
		{
			name:    "arbitrary value with calc",
			classes: "m-[2px] m-[calc(100%-var(--arbitrary))]",
			want:    "m-[calc(100%-var(--arbitrary))]",
		},
		{
			name:    "arbitrary value with length label",
			classes: "m-[2px] m-[length:var(--mystery-var)]",
			want:    "m-[length:var(--mystery-var)]",
		},
		{
			name:    "opacity with arbitrary decimal",
			classes: "opacity-10 opacity-[0.025]",
			want:    "opacity-[0.025]",
		},
		{
			name:    "scale with arbitrary value",
			classes: "scale-75 scale-[1.7]",
			want:    "scale-[1.7]",
		},
		{
			name:    "brightness with arbitrary value",
			classes: "brightness-90 brightness-[1.75]",
			want:    "brightness-[1.75]",
		},
		{
			name:    "min-h arbitrary to zero",
			classes: "min-h-[0.5px] min-h-[0]",
			want:    "min-h-[0]",
		},
		{
			name:    "text size vs text color with labels",
			classes: "text-[0.5px] text-[color:0]",
			want:    "text-[0.5px] text-[color:0]",
		},
		{
			name:    "text size vs text variable",
			classes: "text-[0.5px] text-(--my-0)",
			want:    "text-[0.5px] text-(--my-0)",
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

func TestArbitraryValueWithModifiers(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "hover modifier with arbitrary values",
			classes: "hover:m-[2px] hover:m-[length:var(--c)]",
			want:    "hover:m-[length:var(--c)]",
		},
		{
			name:    "sorted modifier chain with arbitrary values",
			classes: "hover:focus:m-[2px] focus:hover:m-[length:var(--c)]",
			want:    "focus:hover:m-[length:var(--c)]",
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

func TestArbitraryValueBorderColor(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "border width and color label don't conflict",
			classes: "border-b border-[color:rgb(var(--color-gray-500-rgb)/50%))]",
			want:    "border-b border-[color:rgb(var(--color-gray-500-rgb)/50%))]",
		},
		{
			name:    "border color label and width don't conflict reversed",
			classes: "border-[color:rgb(var(--color-gray-500-rgb)/50%))] border-b",
			want:    "border-[color:rgb(var(--color-gray-500-rgb)/50%))] border-b",
		},
		{
			name:    "border color label overridden by color class",
			classes: "border-b border-[color:rgb(var(--color-gray-500-rgb)/50%))] border-some-coloooor",
			want:    "border-b border-some-coloooor",
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

func TestArbitraryValueGridRows(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "grid rows arbitrary to number",
			classes: "grid-rows-[1fr,auto] grid-rows-2",
			want:    "grid-rows-2",
		},
		{
			name:    "grid rows repeat arbitrary to number",
			classes: "grid-rows-[repeat(20,minmax(0,1fr))] grid-rows-3",
			want:    "grid-rows-3",
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

func TestArbitraryValueThemeFunction(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "margin with calc theme",
			classes: "mt-2 mt-[calc(theme(fontSize.4xl)/1.125)]",
			want:    "mt-[calc(theme(fontSize.4xl)/1.125)]",
		},
		{
			name:    "padding with calc theme and multi-value",
			classes: "p-2 p-[calc(theme(fontSize.4xl)/1.125)_10px]",
			want:    "p-[calc(theme(fontSize.4xl)/1.125)_10px]",
		},
		{
			name:    "margin with length label theme",
			classes: "mt-2 mt-[length:theme(someScale.someValue)]",
			want:    "mt-[length:theme(someScale.someValue)]",
		},
		{
			name:    "margin with plain theme",
			classes: "mt-2 mt-[theme(someScale.someValue)]",
			want:    "mt-[theme(someScale.someValue)]",
		},
		{
			name:    "text-size with length label theme",
			classes: "text-2xl text-[length:theme(someScale.someValue)]",
			want:    "text-[length:theme(someScale.someValue)]",
		},
		{
			name:    "text-size with calc theme",
			classes: "text-2xl text-[calc(theme(fontSize.4xl)/1.125)]",
			want:    "text-[calc(theme(fontSize.4xl)/1.125)]",
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

func TestArbitraryValueBgSize(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "bg-cover vs labeled arbitrary bg values",
			classes: "bg-cover bg-[percentage:30%] bg-[size:200px_100px] bg-[length:200px_100px]",
			want:    "bg-[percentage:30%] bg-[length:200px_100px]",
		},
		{
			name:    "bg-none and various bg image types",
			classes: "bg-none bg-[url(.)] bg-[image:.] bg-[url:.] bg-[linear-gradient(.)] bg-linear-to-r",
			want:    "bg-linear-to-r",
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

func TestArbitraryValueFontWeight(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "font number conflicts",
			classes: "font-[400] font-[600]",
			want:    "font-[600]",
		},
		{
			name:    "font var conflicts",
			classes: "font-[var(--a)] font-[var(--b)]",
			want:    "font-[var(--b)]",
		},
		{
			name:    "font weight label vs var",
			classes: "font-[weight:var(--a)] font-[var(--b)]",
			want:    "font-[var(--b)]",
		},
		{
			name:    "font number vs weight label var",
			classes: "font-[400] font-[weight:var(--b)]",
			want:    "font-[weight:var(--b)]",
		},
		{
			name:    "font weight label conflicts",
			classes: "font-[weight:var(--a)] font-[weight:var(--b)]",
			want:    "font-[weight:var(--b)]",
		},
		{
			name:    "font family-name label doesn't conflict with var",
			classes: "font-[family-name:var(--a)] font-[var(--b)]",
			want:    "font-[family-name:var(--a)] font-[var(--b)]",
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

func TestArbitraryVariableValues(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "bg color vs position with variables",
			classes: "bg-red bg-(--other-red) bg-bottom bg-(position:-my-pos)",
			want:    "bg-(--other-red) bg-(position:-my-pos)",
		},
		{
			name:    "shadow with various variable types",
			classes: "shadow-xs shadow-(shadow:--something) shadow-red shadow-(--some-other-shadow) shadow-(color:--some-color)",
			want:    "shadow-(--some-other-shadow) shadow-(color:--some-color)",
		},
		{
			name:    "font variable conflicts",
			classes: "font-(--a) font-(--b)",
			want:    "font-(--b)",
		},
		{
			name:    "font weight variable vs unlabeled variable",
			classes: "font-(weight:--a) font-(--b)",
			want:    "font-(--b)",
		},
		{
			name:    "font family-name variable doesn't conflict with unlabeled",
			classes: "font-(family-name:--a) font-(--b)",
			want:    "font-(family-name:--a) font-(--b)",
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
