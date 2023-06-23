package internal

import (
	"sync"
	"time"
)

type Cache struct {
	cacheAreas map[string]cacheEntry
	mux        sync.RWMutex
}

var intervalTime int
var GlobalCache Cache

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func (c *Cache) add(key string, data []byte) {
	c.cacheAreas[key] = cacheEntry{time.Now(), data}
}

func (c *Cache) get(key string) ([]byte, bool) {
	defer c.mux.Unlock()
	c.mux.Lock()
	data, ok := c.cacheAreas[key]
	if !ok {
		return nil, false
	}
	return data.data, true
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for key, data := range c.cacheAreas {
		if data.createdAt.Before(timeAgo) {
			c.mux.Lock()
			delete(c.cacheAreas, key)
			c.mux.Unlock()
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}
func (c *Cache) NewCache(interval time.Duration) {
	c.cacheAreas = make(map[string]cacheEntry)
	go c.reapLoop(interval)
	// run in a seperate thread

}
