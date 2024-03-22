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
	Mu      *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Entries: make(map[string]cacheEntry),
		Mu:      &sync.RWMutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

// Adds an entry to the Cache
func (c *Cache) Add(key string, val *[]byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       *val,
	}
}

// Searches the Cache for a page
// A cache hit will return the data and true
// A cache miss will return nil, false
func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	entry, ok := c.Entries[key]
	return entry.val, ok
}

// Reaps entries older than Cache Interval after each Interval has passed
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}

}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for name, entry := range c.Entries {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.Entries, name)
		}
	}
}
