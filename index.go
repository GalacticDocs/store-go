package store

import (
	"github.com/GalacticDocs/store-go/cache"
	"github.com/GalacticDocs/store-go/set"
)

func Cache[T any](capacity int) *cache.Cache[T] {
	return cache.New[T](capacity)
}

func Set[T comparable]() *set.Set[T] {
	return set.New[T]()
}