package twmerge

const (
	ImportantModifier = "!"
	modifierSeparator = ':'
)

// CreateParseClassName creates a parser function for the given config.
// The parser splits a class name into modifiers, base class, and metadata.
func CreateParseClassName(config *Config) func(string) ParsedClassName {
	parseClassName := func(className string) ParsedClassName {
		var modifiers []string
		bracketDepth := 0
		parenDepth := 0
		modifierStart := 0
		postfixModifierPosition := -1

		for i := 0; i < len(className); i++ {
			ch := className[i]

			if bracketDepth == 0 && parenDepth == 0 {
				if ch == byte(modifierSeparator) {
					modifiers = append(modifiers, className[modifierStart:i])
					modifierStart = i + 1
					continue
				}

				if ch == '/' {
					postfixModifierPosition = i
					continue
				}
			}

			switch ch {
			case '[':
				bracketDepth++
			case ']':
				bracketDepth--
			case '(':
				parenDepth++
			case ')':
				parenDepth--
			}
		}

		baseClassNameWithImportantModifier := className
		if len(modifiers) > 0 {
			baseClassNameWithImportantModifier = className[modifierStart:]
		}

		baseClassName := baseClassNameWithImportantModifier
		hasImportantModifier := false

		if len(baseClassNameWithImportantModifier) > 0 &&
			baseClassNameWithImportantModifier[len(baseClassNameWithImportantModifier)-1] == '!' {
			baseClassName = baseClassNameWithImportantModifier[:len(baseClassNameWithImportantModifier)-1]
			hasImportantModifier = true
		} else if len(baseClassNameWithImportantModifier) > 0 &&
			baseClassNameWithImportantModifier[0] == '!' {
			// Legacy Tailwind v3 important modifier at start
			baseClassName = baseClassNameWithImportantModifier[1:]
			hasImportantModifier = true
		}

		maybePostfixModifierPosition := -1
		if postfixModifierPosition != -1 && postfixModifierPosition > modifierStart {
			maybePostfixModifierPosition = postfixModifierPosition - modifierStart
		}

		return ParsedClassName{
			Modifiers:                    modifiers,
			HasImportantModifier:         hasImportantModifier,
			BaseClassName:                baseClassName,
			MaybePostfixModifierPosition: maybePostfixModifierPosition,
			IsExternal:                   false,
		}
	}

	if config.Prefix != "" {
		fullPrefix := config.Prefix + string(modifierSeparator)
		originalParseClassName := parseClassName
		parseClassName = func(className string) ParsedClassName {
			if len(className) >= len(fullPrefix) && className[:len(fullPrefix)] == fullPrefix {
				return originalParseClassName(className[len(fullPrefix):])
			}
			return ParsedClassName{
				Modifiers:                    nil,
				HasImportantModifier:         false,
				BaseClassName:                className,
				MaybePostfixModifierPosition: -1,
				IsExternal:                   true,
			}
		}
	}

	return parseClassName
}
