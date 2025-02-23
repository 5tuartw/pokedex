package pokecache

import(
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		mu: sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for i, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, i)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

