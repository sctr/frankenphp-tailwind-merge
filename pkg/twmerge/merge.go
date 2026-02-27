package twmerge

import (
	"strings"
)

var splitClassesRegex = strings.Fields

// MergeClassList merges a space-separated class list, resolving conflicts
// by keeping the last conflicting class (reverse iteration).
func MergeClassList(classList string, utils *ConfigUtils) string {
	classNames := splitClassesRegex(strings.TrimSpace(classList))
	if len(classNames) == 0 {
		return ""
	}

	classGroupsInConflict := make([]string, 0, len(classNames))
	result := ""

	for i := len(classNames) - 1; i >= 0; i-- {
		originalClassName := classNames[i]

		parsed := utils.ParseClassName(originalClassName)

		if parsed.IsExternal {
			result = prepend(originalClassName, result)
			continue
		}

		hasPostfixModifier := parsed.MaybePostfixModifierPosition != -1
		var classGroupID string

		if hasPostfixModifier {
			classGroupID = utils.GetClassGroupID(parsed.BaseClassName[:parsed.MaybePostfixModifierPosition])
		} else {
			classGroupID = utils.GetClassGroupID(parsed.BaseClassName)
		}

		if classGroupID == "" {
			if !hasPostfixModifier {
				// Not a Tailwind class
				result = prepend(originalClassName, result)
				continue
			}

			classGroupID = utils.GetClassGroupID(parsed.BaseClassName)
			if classGroupID == "" {
				// Not a Tailwind class
				result = prepend(originalClassName, result)
				continue
			}
			hasPostfixModifier = false
		}

		// Build variant modifier string
		var variantModifier string
		if len(parsed.Modifiers) == 0 {
			variantModifier = ""
		} else if len(parsed.Modifiers) == 1 {
			variantModifier = parsed.Modifiers[0]
		} else {
			sorted := utils.SortModifiers(parsed.Modifiers)
			variantModifier = strings.Join(sorted, ":")
		}

		modifierID := variantModifier
		if parsed.HasImportantModifier {
			modifierID = variantModifier + ImportantModifier
		}

		classID := modifierID + classGroupID

		if contains(classGroupsInConflict, classID) {
			// Tailwind class omitted due to conflict
			continue
		}

		classGroupsInConflict = append(classGroupsInConflict, classID)

		conflictGroups := utils.GetConflictingClassGroupIDs(classGroupID, hasPostfixModifier)
		for _, group := range conflictGroups {
			classGroupsInConflict = append(classGroupsInConflict, modifierID+group)
		}

		// Tailwind class not in conflict
		result = prepend(originalClassName, result)
	}

	return result
}

func prepend(className, result string) string {
	if result == "" {
		return className
	}
	return className + " " + result
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
