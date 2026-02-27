package twmerge

// ClassValidator is a function that validates whether a class part value
// matches a specific pattern (e.g., is a number, a t-shirt size, etc.).
type ClassValidator func(value string) bool

// ThemeGetter retrieves class group definitions from the theme config.
type ThemeGetter struct {
	Key string
}

// ClassDefinition represents a single entry in a class group definition.
// It can be one of:
//   - string: a literal class part (e.g., "flex", "block")
//   - map[string][]ClassDefinition: nested class parts (e.g., {"overflow": [...]})
//   - ClassValidator: a validator function
//   - ThemeGetter: a reference to theme values
type ClassDefinition interface{}

// Config holds the complete tailwind-merge configuration.
type Config struct {
	CacheSize                      int
	Prefix                         string
	Theme                          map[string][]ClassDefinition
	ClassGroups                    map[string][]ClassDefinition
	ConflictingClassGroups         map[string][]string
	ConflictingClassGroupModifiers map[string][]string
	OrderSensitiveModifiers        []string
}

// ParsedClassName represents a parsed Tailwind CSS class name.
type ParsedClassName struct {
	Modifiers                    []string
	HasImportantModifier         bool
	BaseClassName                string
	MaybePostfixModifierPosition int // -1 means not set
	IsExternal                   bool
}
