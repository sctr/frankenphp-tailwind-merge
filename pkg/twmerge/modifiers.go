package twmerge

import "sort"

// CreateSortModifiers creates a function that sorts modifiers according to:
// - Regular modifiers are sorted alphabetically in segments
// - Arbitrary variants (starting with '[') preserve their position
// - Order-sensitive modifiers preserve their position
func CreateSortModifiers(config *Config) func([]string) []string {
	modifierWeights := make(map[string]bool)
	for _, mod := range config.OrderSensitiveModifiers {
		modifierWeights[mod] = true
	}

	return func(modifiers []string) []string {
		result := make([]string, 0, len(modifiers))
		var currentSegment []string

		for _, modifier := range modifiers {
			isArbitrary := len(modifier) > 0 && modifier[0] == '['
			isOrderSensitive := modifierWeights[modifier]

			if isArbitrary || isOrderSensitive {
				if len(currentSegment) > 0 {
					sort.Strings(currentSegment)
					result = append(result, currentSegment...)
					currentSegment = currentSegment[:0]
				}
				result = append(result, modifier)
			} else {
				currentSegment = append(currentSegment, modifier)
			}
		}

		if len(currentSegment) > 0 {
			sort.Strings(currentSegment)
			result = append(result, currentSegment...)
		}

		return result
	}
}
