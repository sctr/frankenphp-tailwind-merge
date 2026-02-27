package twmerge

import "testing"

func newTestConfig() *Config {
	return &Config{
		Theme: map[string][]ClassDefinition{
			"spacing": {
				IsArbitraryVariable,
				IsArbitraryValue,
				"1", "2", "3", "4", "5", "6", "8", "10", "12", "16", "20", "24",
			},
			"color": {
				IsArbitraryVariable,
				IsArbitraryValue,
				"red", "blue", "green", "white", "black",
				map[string][]ClassDefinition{
					"red": {"100", "200", "300", "400", "500", "600", "700", "800", "900"},
				},
			},
		},
		ClassGroups: map[string][]ClassDefinition{
			"display": {"block", "inline-block", "inline", "flex", "inline-flex", "grid", "hidden", "contents", "flow-root", "list-item"},
			"p": {
				map[string][]ClassDefinition{
					"p": {FromTheme("spacing")},
				},
			},
			"px": {
				map[string][]ClassDefinition{
					"px": {FromTheme("spacing")},
				},
			},
			"py": {
				map[string][]ClassDefinition{
					"py": {FromTheme("spacing")},
				},
			},
			"pt": {
				map[string][]ClassDefinition{
					"pt": {FromTheme("spacing")},
				},
			},
			"pr": {
				map[string][]ClassDefinition{
					"pr": {FromTheme("spacing")},
				},
			},
			"pb": {
				map[string][]ClassDefinition{
					"pb": {FromTheme("spacing")},
				},
			},
			"pl": {
				map[string][]ClassDefinition{
					"pl": {FromTheme("spacing")},
				},
			},
			"m": {
				map[string][]ClassDefinition{
					"m": {FromTheme("spacing")},
				},
			},
			"bg-color": {
				map[string][]ClassDefinition{
					"bg": {FromTheme("color")},
				},
			},
			"text-color": {
				map[string][]ClassDefinition{
					"text": {FromTheme("color")},
				},
			},
			"font-weight": {
				map[string][]ClassDefinition{
					"font": {"thin", "extralight", "light", "normal", "medium", "semibold", "bold", "extrabold", "black"},
				},
			},
			"text-size": {
				map[string][]ClassDefinition{
					"text": {"xs", "sm", "base", "lg", "xl", "2xl", "3xl"},
				},
			},
			"rounded": {
				map[string][]ClassDefinition{
					"rounded": {"", "none", "sm", "md", "lg", "xl", "2xl", "3xl", "full"},
				},
			},
		},
		ConflictingClassGroups: map[string][]string{
			"p":  {"px", "py", "pt", "pr", "pb", "pl"},
			"px": {"pr", "pl"},
			"py": {"pt", "pb"},
		},
		ConflictingClassGroupModifiers: map[string][]string{
			"text-size": {"text-color"},
		},
	}
}

func TestGetClassGroupID_ExactMatch(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	tests := []struct {
		className string
		want      string
	}{
		{"flex", "display"},
		{"block", "display"},
		{"inline", "display"},
		{"hidden", "display"},
		{"grid", "display"},
	}

	for _, tt := range tests {
		got := utils.GetClassGroupID(tt.className)
		if got != tt.want {
			t.Errorf("GetClassGroupID(%q) = %q, want %q", tt.className, got, tt.want)
		}
	}
}

func TestGetClassGroupID_WithValue(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	tests := []struct {
		className string
		want      string
	}{
		{"p-4", "p"},
		{"px-2", "px"},
		{"py-1", "py"},
		{"pt-3", "pt"},
		{"m-4", "m"},
	}

	for _, tt := range tests {
		got := utils.GetClassGroupID(tt.className)
		if got != tt.want {
			t.Errorf("GetClassGroupID(%q) = %q, want %q", tt.className, got, tt.want)
		}
	}
}

func TestGetClassGroupID_ArbitraryValue(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("bg-[#B91C1C]")
	if got != "bg-color" {
		t.Errorf("GetClassGroupID(bg-[#B91C1C]) = %q, want 'bg-color'", got)
	}
}

func TestGetClassGroupID_ArbitraryVariable(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("bg-(--my-color)")
	if got != "bg-color" {
		t.Errorf("GetClassGroupID(bg-(--my-color)) = %q, want 'bg-color'", got)
	}
}

func TestGetClassGroupID_ArbitraryProperty(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("[color:red]")
	if got != "arbitrary..color" {
		t.Errorf("GetClassGroupID([color:red]) = %q, want 'arbitrary..color'", got)
	}
}

func TestGetClassGroupID_ArbitraryPropertyNoColon(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("[something]")
	if got != "" {
		t.Errorf("GetClassGroupID([something]) = %q, want ''", got)
	}
}

func TestGetClassGroupID_Negative(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("-m-4")
	if got != "m" {
		t.Errorf("GetClassGroupID(-m-4) = %q, want 'm'", got)
	}
}

func TestGetClassGroupID_Unknown(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("my-custom-class")
	if got != "" {
		t.Errorf("GetClassGroupID(my-custom-class) = %q, want ''", got)
	}
}

func TestGetClassGroupID_ColorWithShade(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("bg-red-500")
	if got != "bg-color" {
		t.Errorf("GetClassGroupID(bg-red-500) = %q, want 'bg-color'", got)
	}

	got = utils.GetClassGroupID("text-blue")
	if got != "text-color" {
		t.Errorf("GetClassGroupID(text-blue) = %q, want 'text-color'", got)
	}
}

func TestGetClassGroupID_Rounded(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	tests := []struct {
		className string
		want      string
	}{
		{"rounded", "rounded"},
		{"rounded-md", "rounded"},
		{"rounded-full", "rounded"},
	}

	for _, tt := range tests {
		got := utils.GetClassGroupID(tt.className)
		if got != tt.want {
			t.Errorf("GetClassGroupID(%q) = %q, want %q", tt.className, got, tt.want)
		}
	}
}

func TestGetClassGroupID_FontWeight(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	got := utils.GetClassGroupID("font-bold")
	if got != "font-weight" {
		t.Errorf("GetClassGroupID(font-bold) = %q, want 'font-weight'", got)
	}

	got = utils.GetClassGroupID("font-thin")
	if got != "font-weight" {
		t.Errorf("GetClassGroupID(font-thin) = %q, want 'font-weight'", got)
	}
}

func TestGetConflictingIDs(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	conflicts := utils.GetConflictingClassGroupIDs("p", false)
	expected := map[string]bool{"px": true, "py": true, "pt": true, "pr": true, "pb": true, "pl": true}

	if len(conflicts) != len(expected) {
		t.Fatalf("expected %d conflicts, got %d: %v", len(expected), len(conflicts), conflicts)
	}
	for _, c := range conflicts {
		if !expected[c] {
			t.Errorf("unexpected conflict: %q", c)
		}
	}
}

func TestGetConflictingIDs_PostfixModifier(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	// text-size has postfix modifier conflicts with text-color
	conflicts := utils.GetConflictingClassGroupIDs("text-size", true)
	found := false
	for _, c := range conflicts {
		if c == "text-color" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected text-color in conflicts, got %v", conflicts)
	}
}

func TestGetConflictingIDs_NoConflicts(t *testing.T) {
	config := newTestConfig()
	utils := CreateClassGroupUtils(config)

	conflicts := utils.GetConflictingClassGroupIDs("display", false)
	if conflicts != nil {
		t.Errorf("expected no conflicts for display, got %v", conflicts)
	}
}

func TestCreateClassMap_Build(t *testing.T) {
	config := newTestConfig()
	classMap := CreateClassMap(config)

	// Check that the "flex" entry exists in the trie
	flexNode, ok := classMap.NextPart["flex"]
	if !ok {
		t.Fatal("expected 'flex' in class map")
	}
	if flexNode.ClassGroupID != "display" {
		t.Errorf("expected flex classGroupID 'display', got %q", flexNode.ClassGroupID)
	}

	// Check that "p" -> spacing values exist
	pNode, ok := classMap.NextPart["p"]
	if !ok {
		t.Fatal("expected 'p' in class map")
	}
	// "p-4" should resolve via next part "4"
	fourNode, ok := pNode.NextPart["4"]
	if !ok {
		t.Fatal("expected '4' under 'p' in class map")
	}
	if fourNode.ClassGroupID != "p" {
		t.Errorf("expected p-4 classGroupID 'p', got %q", fourNode.ClassGroupID)
	}
}
