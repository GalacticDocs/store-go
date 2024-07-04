package store

import "github.com/GalacticDocs/store-go/cache"

func Cache[T any](capacity int) *cache.Cache[T] {
	return cache.New[T](capacity)
}