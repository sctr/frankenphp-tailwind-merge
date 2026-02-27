package twmerge

import "sync"

// CreateTailwindMerge creates a tailwind merge function with the given config factory.
// The config is lazily initialized on first call.
func CreateTailwindMerge(getConfig func() *Config) func(classes ...string) string {
	var (
		configUtils *ConfigUtils
		once        sync.Once
	)

	return func(classes ...string) string {
		once.Do(func() {
			config := getConfig()
			configUtils = CreateConfigUtils(config)
		})

		classList := TwJoin(classes...)
		if classList == "" {
			return ""
		}

		cached, ok := configUtils.Cache.Get(classList)
		if ok {
			return cached
		}

		result := MergeClassList(classList, configUtils)
		configUtils.Cache.Set(classList, result)
		return result
	}
}

// twMerge is the default TwMerge instance using the default config.
var twMerge func(classes ...string) string
var twMergeOnce sync.Once

// TwMerge merges Tailwind CSS classes using the default configuration.
// This is the main entry point for the library.
func TwMerge(classes ...string) string {
	twMergeOnce.Do(func() {
		twMerge = CreateTailwindMerge(GetDefaultConfig)
	})
	return twMerge(classes...)
}
