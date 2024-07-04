package set

import "sync"

type Set[T comparable] struct {
	data map[T]struct{}
	mu   sync.RWMutex
}
