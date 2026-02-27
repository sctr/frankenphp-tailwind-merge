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
