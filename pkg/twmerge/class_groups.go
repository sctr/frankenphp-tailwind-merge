package twmerge

import (
	"sort"
	"strings"
)

const (
	classPartSeparator     = "-"
	arbitraryPropertyPrefix = "arbitrary.."
)

// ClassPartNode is a trie node for resolving class names to class group IDs.
type ClassPartNode struct {
	NextPart     map[string]*ClassPartNode
	Validators   []ClassValidatorEntry
	ClassGroupID string
}

// ClassValidatorEntry pairs a class group ID with a validator function.
type ClassValidatorEntry struct {
	ClassGroupID string
	Validator    ClassValidator
}

func newClassPartNode() *ClassPartNode {
	return &ClassPartNode{
		NextPart: make(map[string]*ClassPartNode),
	}
}

// ClassGroupUtils provides methods for looking up class group IDs and conflicts.
type ClassGroupUtils struct {
	classMap                       *ClassPartNode
	conflictingClassGroups         map[string][]string
	conflictingClassGroupModifiers map[string][]string
}

// CreateClassGroupUtils builds the trie and conflict lookups from config.
func CreateClassGroupUtils(config *Config) *ClassGroupUtils {
	return &ClassGroupUtils{
		classMap:                       CreateClassMap(config),
		conflictingClassGroups:         config.ConflictingClassGroups,
		conflictingClassGroupModifiers: config.ConflictingClassGroupModifiers,
	}
}

// GetClassGroupID resolves a base class name to its class group ID.
// Returns "" if the class is not recognized.
func (u *ClassGroupUtils) GetClassGroupID(className string) string {
	// Check for arbitrary property: [property:value]
	if len(className) > 1 && className[0] == '[' && className[len(className)-1] == ']' {
		return getGroupIDForArbitraryProperty(className)
	}

	parts := strings.Split(className, classPartSeparator)
	startIndex := 0
	// Handle negative values like "-m-4" where parts[0] == ""
	if parts[0] == "" && len(parts) > 1 {
		startIndex = 1
	}
	return getGroupRecursive(parts, startIndex, u.classMap)
}

// GetConflictingClassGroupIDs returns all class group IDs that conflict with
// the given class group ID.
func (u *ClassGroupUtils) GetConflictingClassGroupIDs(classGroupID string, hasPostfixModifier bool) []string {
	if hasPostfixModifier {
		modConflicts := u.conflictingClassGroupModifiers[classGroupID]
		baseConflicts := u.conflictingClassGroups[classGroupID]

		if modConflicts != nil {
			if baseConflicts != nil {
				result := make([]string, len(baseConflicts)+len(modConflicts))
				copy(result, baseConflicts)
				copy(result[len(baseConflicts):], modConflicts)
				return result
			}
			return modConflicts
		}
		if baseConflicts != nil {
			return baseConflicts
		}
		return nil
	}

	return u.conflictingClassGroups[classGroupID]
}

func getGroupRecursive(classParts []string, startIndex int, node *ClassPartNode) string {
	remaining := len(classParts) - startIndex
	if remaining == 0 {
		return node.ClassGroupID
	}

	currentPart := classParts[startIndex]
	if nextNode, ok := node.NextPart[currentPart]; ok {
		result := getGroupRecursive(classParts, startIndex+1, nextNode)
		if result != "" {
			return result
		}
	}

	if len(node.Validators) == 0 {
		return ""
	}

	// Build the class rest for validator checking
	var classRest string
	if startIndex == 0 {
		classRest = strings.Join(classParts, classPartSeparator)
	} else {
		classRest = strings.Join(classParts[startIndex:], classPartSeparator)
	}

	for _, v := range node.Validators {
		if v.Validator(classRest) {
			return v.ClassGroupID
		}
	}

	return ""
}

func getGroupIDForArbitraryProperty(className string) string {
	content := className[1 : len(className)-1]
	colonIdx := strings.Index(content, ":")
	if colonIdx == -1 {
		return ""
	}
	property := content[:colonIdx]
	if property == "" {
		return ""
	}
	return arbitraryPropertyPrefix + property
}

// CreateClassMap builds a trie of class part nodes from the config.
func CreateClassMap(config *Config) *ClassPartNode {
	classMap := newClassPartNode()

	// Sort class group keys to ensure deterministic validator ordering on
	// shared trie nodes. Shorter keys first ensures base groups (e.g. "shadow")
	// register their specific validators before derived groups (e.g.
	// "shadow-color") register catch-all validators. This matches the JS
	// reference implementation which relies on object insertion order.
	keys := make([]string, 0, len(config.ClassGroups))
	for k := range config.ClassGroups {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if len(keys[i]) != len(keys[j]) {
			return len(keys[i]) < len(keys[j])
		}
		return keys[i] < keys[j]
	})

	for _, classGroupID := range keys {
		processClassesRecursively(config.ClassGroups[classGroupID], classMap, classGroupID, config.Theme)
	}

	return classMap
}

func processClassesRecursively(classGroup []ClassDefinition, node *ClassPartNode, classGroupID string, theme map[string][]ClassDefinition) {
	for _, def := range classGroup {
		processClassDefinition(def, node, classGroupID, theme)
	}
}

func processClassDefinition(def ClassDefinition, node *ClassPartNode, classGroupID string, theme map[string][]ClassDefinition) {
	switch v := def.(type) {
	case string:
		processStringDefinition(v, node, classGroupID)
	case ClassValidator:
		node.Validators = append(node.Validators, ClassValidatorEntry{
			ClassGroupID: classGroupID,
			Validator:    v,
		})
	case func(string) bool:
		node.Validators = append(node.Validators, ClassValidatorEntry{
			ClassGroupID: classGroupID,
			Validator:    v,
		})
	case ThemeGetter:
		themeGroup := theme[v.Key]
		if themeGroup != nil {
			processClassesRecursively(themeGroup, node, classGroupID, theme)
		}
	case map[string][]ClassDefinition:
		for key, value := range v {
			subNode := getPart(node, key)
			processClassesRecursively(value, subNode, classGroupID, theme)
		}
	}
}

func processStringDefinition(classDef string, node *ClassPartNode, classGroupID string) {
	if classDef == "" {
		node.ClassGroupID = classGroupID
		return
	}
	target := getPart(node, classDef)
	target.ClassGroupID = classGroupID
}

func getPart(node *ClassPartNode, path string) *ClassPartNode {
	current := node
	parts := strings.Split(path, classPartSeparator)
	for _, part := range parts {
		next, ok := current.NextPart[part]
		if !ok {
			next = newClassPartNode()
			current.NextPart[part] = next
		}
		current = next
	}
	return current
}
