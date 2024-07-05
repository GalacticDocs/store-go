package weakset

import (
	"sync"
)

type weakValue[T comparable] struct {
	refs []*T
}

type WeakSet[T comparable] struct {
	data map[T]*weakValue[T]
	mu sync.RWMutex
}