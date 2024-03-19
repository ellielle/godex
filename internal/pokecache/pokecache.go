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
	Entries  map[string]cacheEntry
	Mu       *sync.RWMutex
	Interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Entries:  make(map[string]cacheEntry),
		Mu:       &sync.RWMutex{},
		Interval: interval,
	}
	go cache.reapLoop(interval)

	return cache
}

// Adds an entry to the Cache
func (ca *Cache) Add(key string, val []byte) {
	ca.Mu.Lock()
	defer ca.Mu.Unlock()

	ca.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Searches the Cache for a page
// A cache hit will return the data and true
// A cache miss will return nil, false
func (ca *Cache) Get(key string) ([]byte, bool) {
	ca.Mu.Lock()
	defer ca.Mu.Unlock()

	entry, ok := ca.Entries[key]
	return entry.val, ok
}

// Reaps entries older than Cache Interval after each Interval has passed
func (ca *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	<-ticker.C
	for name, entry := range ca.Entries {
		if entry.createdAt.Before(time.Now().Add(ca.Interval)) {
			delete(ca.Entries, name)
		}
	}
}
