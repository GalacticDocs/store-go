package cache

import (
	"time"

	storing_errors "github.com/GalacticDocs/store-go/errors"
)

// Initialises a new Cache instance.
func New[T any](capacity int) *Cache[T] {
	return &Cache[T]{
		data:     make(map[string]cacheData[T]),
		capacity: capacity,
	}
}

// Removes any value associated to the key.
func (c *Cache[T]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[key]
	if !ok {
		storing_errors.NoKeyError("delete", key)

		return
	}

	delete(c.data, key)
}

// Returns the value associated to the key, or nil if there is none.
func (c *Cache[T]) Fetch(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var (
		payload cacheData[T]
		exists  bool
	)

	for k, data := range c.data {
		if k == key {
			payload = data
			exists = true
		} else {
			exists = false
		}
	}

	if !exists {
		storing_errors.KeyNotFoundError(key, "get")

		return *new(T), false
	}

	if time.Now().After(payload.expiration) {
		delete(c.data, key)

		return *new(T), false
	}

	return payload.value, true
}

// Returns the value associated to the key, or nil if there is none. NOTE: for stricter fetching you may want to consider using Cache().Fetch(key)
func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, ok := c.data[key]
	if !ok {
		storing_errors.KeyNotFoundError(key, "get")

		return *new(T), false
	}

	if time.Now().After(data.expiration) {
		delete(c.data, key)

		return *new(T), false
	}

	return data.value, true
}

func (c *Cache[T]) Set(key string, value T, expiration time.Time) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !storing_errors.CheckIfBadKey(key) {
		return false
	}

	if key == "" {
		storing_errors.NoKeyError("set", key)

		return false
	}

	if len(c.data) >= c.capacity {
		c.evictExpired()
	}

	c.data[key] = cacheData[T]{
		value:      value,
		expiration: expiration,
	}

	return true
}

func (c *Cache[T]) evictExpired() {
	for key, data := range c.data {
		if time.Now().After(data.expiration) {
			delete(c.data, key)
		}
	}
}
