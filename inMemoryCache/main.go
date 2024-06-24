package main

import (
	"fmt"
	"sync"
	"time"
)

var _ Cache = (*InMemoryCache)(nil)

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	expireIn time.Duration
	cache    map[string]CacheEntry
	mutex    sync.RWMutex
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn,
		cache:    make(map[string]CacheEntry),
	}
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}
}

func (c *InMemoryCache) Get(key string) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil
	}

	if time.Since(entry.settledAt) > c.expireIn {
		delete(c.cache, key)
		return nil
	}

	return entry.value
}

func main() {
	cache := NewInMemoryCache(5 * time.Minute)

	cache.Set("key1", "value1")

	value := cache.Get("key1")
	fmt.Println(value)

	time.Sleep(6 * time.Minute)
	value = cache.Get("key1")
	fmt.Println(value)
}
