package twmerge

import (
	"reflect"
	"testing"
)

func TestSortModifiers_Alphabetical(t *testing.T) {
	sortMods := CreateSortModifiers(&Config{})

	result := sortMods([]string{"hover", "dark", "focus"})
	expected := []string{"dark", "focus", "hover"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSortModifiers_AlreadySorted(t *testing.T) {
	sortMods := CreateSortModifiers(&Config{})

	result := sortMods([]string{"dark", "focus", "hover"})
	expected := []string{"dark", "focus", "hover"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSortModifiers_SingleModifier(t *testing.T) {
	sortMods := CreateSortModifiers(&Config{})

	result := sortMods([]string{"hover"})
	expected := []string{"hover"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSortModifiers_Empty(t *testing.T) {
	sortMods := CreateSortModifiers(&Config{})

	result := sortMods([]string{})
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func TestSortModifiers_ArbitraryVariant(t *testing.T) {
	sortMods := CreateSortModifiers(&Config{})

	result := sortMods([]string{"hover", "[&>svg]", "focus"})
	expected := []string{"hover", "[&>svg]", "focus"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSortModifiers_OrderSensitive(t *testing.T) {
	sortMods := CreateSortModifiers(&Config{
		OrderSensitiveModifiers: []string{"before", "after"},
	})

	result := sortMods([]string{"hover", "before", "focus"})
	// "hover" is sorted, then "before" preserved, then "focus"
	expected := []string{"hover", "before", "focus"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSortModifiers_MixedSegments(t *testing.T) {
	sortMods := CreateSortModifiers(&Config{
		OrderSensitiveModifiers: []string{"before"},
	})

	result := sortMods([]string{"hover", "dark", "before", "focus", "active"})
	// Segment 1: ["hover", "dark"] → sorted: ["dark", "hover"]
	// Then "before" (order-sensitive)
	// Segment 2: ["focus", "active"] → sorted: ["active", "focus"]
	expected := []string{"dark", "hover", "before", "active", "focus"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Merge-level modifier tests (from Node.js tests/modifiers.test.ts)

func TestMergeWithPrefixModifiers(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "hover conflict",
			classes: "hover:block hover:inline",
			want:    "hover:inline",
		},
		{
			name:    "different modifier chains don't conflict",
			classes: "hover:block hover:focus:inline",
			want:    "hover:block hover:focus:inline",
		},
		{
			name:    "reordered modifier chains conflict",
			classes: "hover:block hover:focus:inline focus:hover:inline",
			want:    "hover:block focus:hover:inline",
		},
		{
			name:    "focus-within conflict",
			classes: "focus-within:inline focus-within:block",
			want:    "focus-within:block",
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

func TestMergeWithPostfixModifiers(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "same base different postfix",
			classes: "text-lg/7 text-lg/8",
			want:    "text-lg/8",
		},
		{
			name:    "postfix and leading don't conflict",
			classes: "text-lg/none leading-9",
			want:    "text-lg/none leading-9",
		},
		{
			name:    "leading overridden by postfix",
			classes: "leading-9 text-lg/none",
			want:    "text-lg/none",
		},
		{
			name:    "w-full overridden by w-1/2 (fraction not postfix)",
			classes: "w-full w-1/2",
			want:    "w-1/2",
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

func TestMergeModifierSorting(t *testing.T) {
	tests := []struct {
		name    string
		classes string
		want    string
	}{
		{
			name:    "sorted modifiers conflict",
			classes: "c:d:e:block d:c:e:inline",
			want:    "d:c:e:inline",
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
