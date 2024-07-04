package cache

import (
	"sync"
	"time"
)

type cacheData[T any] struct {
	value      T
	expiration time.Time
}

type Cache[T any] struct {
	data     map[string]cacheData[T]
	mu       sync.RWMutex
	capacity int
}
