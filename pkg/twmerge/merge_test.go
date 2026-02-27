package twmerge

import "testing"

func TestTwMerge(t *testing.T) {
	tests := []struct {
		name    string
		classes []string
		want    string
	}{
		{
			name:    "empty input",
			classes: []string{},
			want:    "",
		},
		{
			name:    "single class",
			classes: []string{"px-2"},
			want:    "px-2",
		},
		{
			name:    "empty strings",
			classes: []string{"", "", ""},
			want:    "",
		},
		{
			name:    "shorthand overrides longhand",
			classes: []string{"px-2 py-1", "p-3"},
			want:    "p-3",
		},
		{
			name:    "last conflicting wins",
			classes: []string{"text-red-500", "text-blue-500"},
			want:    "text-blue-500",
		},
		{
			name:    "modifier scoping",
			classes: []string{"hover:bg-red-500", "hover:bg-blue-500", "bg-green-500"},
			want:    "hover:bg-blue-500 bg-green-500",
		},
		{
			name:    "arbitrary values",
			classes: []string{"bg-red-500", "bg-[#1da1f2]"},
			want:    "bg-[#1da1f2]",
		},
		{
			name:    "non-TW preserved",
			classes: []string{"my-custom-class px-2", "px-4"},
			want:    "my-custom-class px-4",
		},
		{
			name:    "important modifier at start",
			classes: []string{"!font-bold", "!font-thin"},
			want:    "!font-thin",
		},
		{
			name:    "important modifier at end",
			classes: []string{"font-bold!", "font-thin!"},
			want:    "font-thin!",
		},
		{
			name:    "postfix modifiers",
			classes: []string{"text-lg/7", "text-lg/8"},
			want:    "text-lg/8",
		},
		{
			name:    "readme hero example",
			classes: []string{"px-2 py-1 bg-red hover:bg-dark-red", "p-3 bg-[#B91C1C]"},
			want:    "hover:bg-dark-red p-3 bg-[#B91C1C]",
		},
		{
			name:    "component override",
			classes: []string{"inline-flex items-center px-4 py-2 bg-blue-600 text-white font-medium rounded-md", "bg-red-600 py-3"},
			want:    "inline-flex items-center px-4 text-white font-medium rounded-md bg-red-600 py-3",
		},
		{
			name:    "display conflicts",
			classes: []string{"block", "hidden"},
			want:    "hidden",
		},
		{
			name:    "flex direction",
			classes: []string{"flex-row", "flex-col"},
			want:    "flex-col",
		},
		{
			name:    "border radius",
			classes: []string{"rounded-md", "rounded-lg"},
			want:    "rounded-lg",
		},
		{
			name:    "position",
			classes: []string{"relative", "absolute"},
			want:    "absolute",
		},
		{
			name:    "different groups don't conflict",
			classes: []string{"p-4 m-4"},
			want:    "p-4 m-4",
		},
		{
			name:    "arbitrary property",
			classes: []string{"[color:red]", "[color:blue]"},
			want:    "[color:blue]",
		},
		{
			name:    "negative margin",
			classes: []string{"-m-4", "-m-8"},
			want:    "-m-8",
		},
		{
			name:    "multiple separate arguments",
			classes: []string{"px-2", "py-1", "p-3"},
			want:    "p-3",
		},
		{
			name:    "width values",
			classes: []string{"w-full", "w-1/2"},
			want:    "w-1/2",
		},
		{
			name:    "gap values",
			classes: []string{"gap-4", "gap-8"},
			want:    "gap-8",
		},
		{
			name:    "text size",
			classes: []string{"text-sm", "text-lg"},
			want:    "text-lg",
		},
		{
			name:    "opacity",
			classes: []string{"opacity-50", "opacity-100"},
			want:    "opacity-100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwMerge(tt.classes...)
			if got != tt.want {
				t.Errorf("TwMerge(%v) = %q, want %q", tt.classes, got, tt.want)
			}
		})
	}
}

func TestTwMerge_CachingWorks(t *testing.T) {
	// Call twice with same input — second should hit cache
	result1 := TwMerge("px-2 py-1", "p-3")
	result2 := TwMerge("px-2 py-1", "p-3")

	if result1 != result2 {
		t.Errorf("cache inconsistency: %q vs %q", result1, result2)
	}
	if result1 != "p-3" {
		t.Errorf("expected 'p-3', got %q", result1)
	}
}

func TestCreateTailwindMerge_CustomConfig(t *testing.T) {
	merge := CreateTailwindMerge(func() *Config {
		return &Config{
			CacheSize: 10,
			ClassGroups: map[string][]ClassDefinition{
				"display": {"block", "hidden", "flex"},
				"color":   {m("text", d("red", "blue", "green")...)},
			},
			ConflictingClassGroups: map[string][]string{},
		}
	})

	result := merge("block", "hidden")
	if result != "hidden" {
		t.Errorf("expected 'hidden', got %q", result)
	}

	result = merge("text-red", "text-blue")
	if result != "text-blue" {
		t.Errorf("expected 'text-blue', got %q", result)
	}
}
