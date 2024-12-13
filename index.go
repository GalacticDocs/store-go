package store

import (
	cache_storage "github.com/GalacticDocs/store-go/cache"
	set_storage "github.com/GalacticDocs/store-go/set"
	map_storage "github.com/GalacticDocs/store-go/map"
	collection_storage "github.com/GalacticDocs/store-go/collection"
	weakset_storage "github.com/GalacticDocs/store-go/weakset"
)

func Cache[T any](capacity int) *cache.Cache[T] {
	return cache_storage.New[T](capacity)
}

func WeakSet[T comparable]() *weakset_storage.WeakSet[T] {
	return weakset_storage.New[T]()
}

func Set[T comparable]() *set.Set[T] {
	return set_storage.New[T]()
}

func Map() *map_storage.IMap {
	return map_storage.Map()
}

func Collection() *collection_storage.ICollection {
	return collection_storage.New()
}
