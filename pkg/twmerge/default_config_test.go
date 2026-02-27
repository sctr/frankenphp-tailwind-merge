package twmerge

import "testing"

func TestDefaultConfig_HasAllPaddingGroups(t *testing.T) {
	config := GetDefaultConfig()

	paddingGroups := []string{"p", "px", "py", "ps", "pe", "pbs", "pbe", "pt", "pr", "pb", "pl"}
	for _, g := range paddingGroups {
		if _, ok := config.ClassGroups[g]; !ok {
			t.Errorf("missing padding group: %s", g)
		}
	}
}

func TestDefaultConfig_HasAllMarginGroups(t *testing.T) {
	config := GetDefaultConfig()

	marginGroups := []string{"m", "mx", "my", "ms", "me", "mbs", "mbe", "mt", "mr", "mb", "ml"}
	for _, g := range marginGroups {
		if _, ok := config.ClassGroups[g]; !ok {
			t.Errorf("missing margin group: %s", g)
		}
	}
}

func TestDefaultConfig_HasDisplayGroup(t *testing.T) {
	config := GetDefaultConfig()

	if _, ok := config.ClassGroups["display"]; !ok {
		t.Fatal("missing display class group")
	}
}

func TestDefaultConfig_HasColorGroups(t *testing.T) {
	config := GetDefaultConfig()

	colorGroups := []string{"bg-color", "text-color", "border-color"}
	for _, g := range colorGroups {
		if _, ok := config.ClassGroups[g]; !ok {
			t.Errorf("missing color group: %s", g)
		}
	}
}

func TestDefaultConfig_ConflictMaps(t *testing.T) {
	config := GetDefaultConfig()

	pConflicts := config.ConflictingClassGroups["p"]
	if pConflicts == nil {
		t.Fatal("missing p conflicts")
	}

	expected := map[string]bool{
		"px": true, "py": true, "ps": true, "pe": true,
		"pbs": true, "pbe": true, "pt": true, "pr": true, "pb": true, "pl": true,
	}
	for _, c := range pConflicts {
		if !expected[c] {
			t.Errorf("unexpected p conflict: %s", c)
		}
	}
	if len(pConflicts) != len(expected) {
		t.Errorf("expected %d p conflicts, got %d", len(expected), len(pConflicts))
	}
}

func TestDefaultConfig_OrderSensitiveModifiers(t *testing.T) {
	config := GetDefaultConfig()

	if len(config.OrderSensitiveModifiers) == 0 {
		t.Fatal("expected non-empty order sensitive modifiers")
	}

	expectedModifiers := map[string]bool{
		"before": true, "after": true, "placeholder": true,
		"file": true, "marker": true, "selection": true,
	}
	found := 0
	for _, mod := range config.OrderSensitiveModifiers {
		if expectedModifiers[mod] {
			found++
		}
	}
	if found != len(expectedModifiers) {
		t.Errorf("missing expected order sensitive modifiers, found %d of %d", found, len(expectedModifiers))
	}
}

func TestDefaultConfig_CacheSize(t *testing.T) {
	config := GetDefaultConfig()
	if config.CacheSize != 500 {
		t.Errorf("expected cache size 500, got %d", config.CacheSize)
	}
}

func TestDefaultConfig_ThemeHasExpectedKeys(t *testing.T) {
	config := GetDefaultConfig()

	expectedThemeKeys := []string{
		"color", "spacing", "blur", "radius", "shadow",
		"animate", "font", "font-weight", "text", "leading", "tracking",
	}
	for _, key := range expectedThemeKeys {
		if _, ok := config.Theme[key]; !ok {
			t.Errorf("missing theme key: %s", key)
		}
	}
}
