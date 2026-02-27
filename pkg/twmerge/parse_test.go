package twmerge

import (
	"reflect"
	"testing"
)

func TestParse_SimpleClass(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("px-4")

	if result.BaseClassName != "px-4" {
		t.Errorf("expected base class 'px-4', got %q", result.BaseClassName)
	}
	if len(result.Modifiers) != 0 {
		t.Errorf("expected no modifiers, got %v", result.Modifiers)
	}
	if result.HasImportantModifier {
		t.Error("expected no important modifier")
	}
	if result.IsExternal {
		t.Error("expected not external")
	}
}

func TestParse_WithModifier(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("hover:bg-red-500")

	if result.BaseClassName != "bg-red-500" {
		t.Errorf("expected base class 'bg-red-500', got %q", result.BaseClassName)
	}
	if !reflect.DeepEqual(result.Modifiers, []string{"hover"}) {
		t.Errorf("expected [hover], got %v", result.Modifiers)
	}
}

func TestParse_MultipleModifiers(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("dark:hover:bg-red-500")

	if result.BaseClassName != "bg-red-500" {
		t.Errorf("expected base class 'bg-red-500', got %q", result.BaseClassName)
	}
	if !reflect.DeepEqual(result.Modifiers, []string{"dark", "hover"}) {
		t.Errorf("expected [dark hover], got %v", result.Modifiers)
	}
}

func TestParse_ImportantModifierEnd(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("font-bold!")

	if result.BaseClassName != "font-bold" {
		t.Errorf("expected base class 'font-bold', got %q", result.BaseClassName)
	}
	if !result.HasImportantModifier {
		t.Error("expected important modifier")
	}
}

func TestParse_ImportantModifierStart(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("!font-bold")

	if result.BaseClassName != "font-bold" {
		t.Errorf("expected base class 'font-bold', got %q", result.BaseClassName)
	}
	if !result.HasImportantModifier {
		t.Error("expected important modifier")
	}
}

func TestParse_PostfixModifier(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("text-lg/7")

	if result.BaseClassName != "text-lg/7" {
		t.Errorf("expected base class 'text-lg/7', got %q", result.BaseClassName)
	}
	if result.MaybePostfixModifierPosition != 7 {
		t.Errorf("expected postfix modifier position 7, got %d", result.MaybePostfixModifierPosition)
	}
}

func TestParse_ArbitraryValue(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("bg-[#B91C1C]")

	if result.BaseClassName != "bg-[#B91C1C]" {
		t.Errorf("expected base class 'bg-[#B91C1C]', got %q", result.BaseClassName)
	}
	// Colons inside brackets should NOT be treated as modifier separators
	if len(result.Modifiers) != 0 {
		t.Errorf("expected no modifiers, got %v", result.Modifiers)
	}
}

func TestParse_ArbitraryValueWithColon(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("bg-[color:red]")

	if result.BaseClassName != "bg-[color:red]" {
		t.Errorf("expected base class 'bg-[color:red]', got %q", result.BaseClassName)
	}
	if len(result.Modifiers) != 0 {
		t.Errorf("expected no modifiers (colon inside brackets), got %v", result.Modifiers)
	}
}

func TestParse_WithPrefix(t *testing.T) {
	parse := CreateParseClassName(&Config{Prefix: "tw"})

	// Class with matching prefix
	result := parse("tw:px-4")
	if result.BaseClassName != "px-4" {
		t.Errorf("expected base class 'px-4', got %q", result.BaseClassName)
	}
	if result.IsExternal {
		t.Error("expected not external for prefixed class")
	}

	// Class without prefix → external
	result = parse("px-4")
	if result.BaseClassName != "px-4" {
		t.Errorf("expected base class 'px-4', got %q", result.BaseClassName)
	}
	if !result.IsExternal {
		t.Error("expected external for non-prefixed class")
	}
}

func TestParse_ModifierWithPrefix(t *testing.T) {
	parse := CreateParseClassName(&Config{Prefix: "tw"})
	result := parse("tw:hover:bg-red-500")

	if result.BaseClassName != "bg-red-500" {
		t.Errorf("expected base class 'bg-red-500', got %q", result.BaseClassName)
	}
	if !reflect.DeepEqual(result.Modifiers, []string{"hover"}) {
		t.Errorf("expected [hover], got %v", result.Modifiers)
	}
}

func TestParse_ArbitraryVariable(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("bg-(--my-color)")

	if result.BaseClassName != "bg-(--my-color)" {
		t.Errorf("expected base class 'bg-(--my-color)', got %q", result.BaseClassName)
	}
	// Colons inside parens should NOT be treated as modifier separators
	if len(result.Modifiers) != 0 {
		t.Errorf("expected no modifiers, got %v", result.Modifiers)
	}
}

func TestParse_PostfixNotSet(t *testing.T) {
	parse := CreateParseClassName(&Config{})
	result := parse("px-4")

	if result.MaybePostfixModifierPosition != -1 {
		t.Errorf("expected postfix position -1, got %d", result.MaybePostfixModifierPosition)
	}
}
