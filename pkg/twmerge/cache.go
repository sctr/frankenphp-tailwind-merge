package twmerge

import "sync"

// LRUCache is a two-tier LRU cache for memoizing merge results.
// When the main cache exceeds maxSize, it rotates to previousCache
// and creates a fresh main cache. Gets check main first, then previous
// with promotion back to main on hit.
type LRUCache struct {
	mu            sync.Mutex
	cache         map[string]string
	previousCache map[string]string
	maxSize       int
}

// NewLRUCache creates a new two-tier LRU cache with the given max size.
// If maxSize < 1, returns a no-op cache.
func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		cache:         make(map[string]string),
		previousCache: make(map[string]string),
		maxSize:       maxSize,
	}
}

// Get retrieves a value from the cache. Returns the value and true if found,
// or empty string and false if not found.
func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.maxSize < 1 {
		return "", false
	}

	if val, ok := c.cache[key]; ok {
		return val, true
	}

	if val, ok := c.previousCache[key]; ok {
		// Promote to main cache
		c.update(key, val)
		return val, true
	}

	return "", false
}

// Set stores a value in the cache.
func (c *LRUCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.maxSize < 1 {
		return
	}

	c.update(key, value)
}

func (c *LRUCache) update(key, value string) {
	c.cache[key] = value

	if len(c.cache) > c.maxSize {
		c.previousCache = c.cache
		c.cache = make(map[string]string)
	}
}
