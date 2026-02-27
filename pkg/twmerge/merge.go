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

	classGroupsInConflict := make(map[string]struct{})
	finalClasses := make([]string, len(classNames))
	cursor := len(classNames)

	for i := len(classNames) - 1; i >= 0; i-- {
		originalClassName := classNames[i]

		parsed := utils.ParseClassName(originalClassName)

		if parsed.IsExternal {
			cursor--
			finalClasses[cursor] = originalClassName
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
				cursor--
				finalClasses[cursor] = originalClassName
				continue
			}

			classGroupID = utils.GetClassGroupID(parsed.BaseClassName)
			if classGroupID == "" {
				cursor--
				finalClasses[cursor] = originalClassName
				continue
			}
			hasPostfixModifier = false
		}

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

		if _, exists := classGroupsInConflict[classID]; exists {
			continue
		}

		classGroupsInConflict[classID] = struct{}{}

		conflictGroups := utils.GetConflictingClassGroupIDs(classGroupID, hasPostfixModifier)
		for _, group := range conflictGroups {
			classGroupsInConflict[modifierID+group] = struct{}{}
		}

		cursor--
		finalClasses[cursor] = originalClassName
	}

	return strings.Join(finalClasses[cursor:], " ")
}
