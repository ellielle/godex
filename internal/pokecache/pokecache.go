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
	Entries map[string]cacheEntry
	mu      *sync.RWMutex
}

// Create a new cache for the session
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

// Adds an entry to the Cache
func (c *Cache) Add(key string, val *[]byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       *val,
	}
}

// Searches the Cache for a page
// A cache hit will return the data and true
// A cache miss will return nil, false
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.Entries[key]
	return entry.val, ok
}

// Controls the loop for reaping old cache entries
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}

}

// Reaps entries older than Cache Interval after each Interval has passed
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for name, entry := range c.Entries {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.Entries, name)
		}
	}
}
