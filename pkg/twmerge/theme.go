package twmerge

// FromTheme creates a ThemeGetter that retrieves class definitions from the
// theme config under the given key.
func FromTheme(key string) ThemeGetter {
	return ThemeGetter{Key: key}
}
