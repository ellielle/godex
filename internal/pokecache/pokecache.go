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

type handleCache interface {
	setCacheInterval(interval time.Duration) *Cache
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
}

func NewCache(interval time.Duration) *Cache {
	var cacheMap map[string]cacheEntry
	cache := &Cache{
		Entries: cacheMap,
		Mu:      &sync.RWMutex{},
	}
	cache.setCacheInterval(interval)

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
	if !ok {
		return nil, false
	}
	return entry.val, true
}

// When triggered, it will remove entries older than the Cache interval
func (ca *Cache) reapLoop() {}

func (ca *Cache) setCacheInterval(interval time.Duration) {
	ca.Mu.Lock()
	defer ca.Mu.Unlock()
	ca.Interval = interval
}
