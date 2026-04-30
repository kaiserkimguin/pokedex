// Package pokecache: this Package is designed to Setup the Cache in such a way, that Map entries
// don't have to be reloaded consantly. The cache is checked on a time basis
// and protected by a mutex. Functions that will use this Cache are 'Map'.
package pokecache

import (
	// sync for mutexes and time to make the cache check time based.
	"sync"
	"time"
)

type Cache struct {
	entries    map[string]CacheEntry
	cacheMutex sync.Mutex
}
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:    map[string]CacheEntry{},
		cacheMutex: sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()
	val, ok := c.entries[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.cacheMutex.Lock()
		for key, value := range c.entries {
			if time.Since(value.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.cacheMutex.Unlock()
	}
}
