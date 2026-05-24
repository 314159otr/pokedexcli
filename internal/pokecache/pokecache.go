package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time	
	val 	  []byte
}

type Cache struct {
	cache 	 map[string]cacheEntry
	mutex 	 sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache {
		cache: make(map[string]cacheEntry),
		mutex: sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cacheEntry, found := cache.cache[key]
	if !found {
		return nil, false
	}
	return cacheEntry.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(interval)
	}
}

func (cache *Cache) reap(interval time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	for key, cacheEntry := range cache.cache {
		elapsed := time.Since(cacheEntry.createdAt)
		if elapsed > interval {
			delete(cache.cache, key)
		}
	}

}
