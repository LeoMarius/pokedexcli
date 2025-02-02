package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.RWMutex
	entries  map[string]cacheEntry
	interval time.Duration
}

// NewCache creates a new Cache with a reap interval
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

// Add inserts a new value into the cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, found := c.entries[key]
	if !found {
		return nil, false
	}
	return entry.val, true
}

// reapLoop periodically removes expired cache entries
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
