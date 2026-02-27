package twmerge

// ConfigUtils holds all the utilities needed for class merging,
// wired together from a single Config.
type ConfigUtils struct {
	Cache                      *LRUCache
	ParseClassName             func(string) ParsedClassName
	SortModifiers              func([]string) []string
	GetClassGroupID            func(string) string
	GetConflictingClassGroupIDs func(string, bool) []string
}

// CreateConfigUtils creates all utilities from the given config.
func CreateConfigUtils(config *Config) *ConfigUtils {
	cache := NewLRUCache(config.CacheSize)
	parseClassName := CreateParseClassName(config)
	sortModifiers := CreateSortModifiers(config)
	classGroupUtils := CreateClassGroupUtils(config)

	return &ConfigUtils{
		Cache:                       cache,
		ParseClassName:              parseClassName,
		SortModifiers:               sortModifiers,
		GetClassGroupID:             classGroupUtils.GetClassGroupID,
		GetConflictingClassGroupIDs: classGroupUtils.GetConflictingClassGroupIDs,
	}
}
