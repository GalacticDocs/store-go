package store

import (
	storing_errors "github.com/GalacticDocs/store-go/errors"
)

type IMapAll struct {
	Key   string
	Value any
}

type IMap struct {
	// Returns all the existing keys in the Map.
	All func() []IMapAll
	// The Clear() method removes all elements from the Map. Returns whether the map has been cleared or not.
	Clear func() bool
	// Removes any value associated to the key. Map().has(key) will return false afterwards.
	Delete func(key string) bool
	// Returns a Boolean asserting whether a value has been associated to the key in the Map or not.
	Exists func(key string) bool
	// Returns the value associated to the key, or nil if there is none.
	Fetch func(key string) any
	// Returns the value associated to the key, or nil if there is none. NOTE: for stricter fetching you may want to consider using Map().fetch(key)
	Get func(key string) any
	// Returns a Boolean asserting whether a value has been associated to the key in the Map or not.
	Has func(key string) bool
	// The iterator which is used to iterate through the map.
	Iterator func() map[string]any
	// Returns all the keys as an array.
	Keys func() []string
	// Sets the value for the key in the Map. Returns whether the key value has been set.
	Set func(key string, value any) bool
	// Returns the amount of keys in the Map.
	Size func() int
	// Updates the value for the provided key in the Map. Returns whether the key value has been changed.
	Update func(key string, value any) bool
	// Returns all the values as an array.
	Values func() []any
}

// Initiates a new Map instance.
func Map() *IMap {
	var (
		store  = make(map[string]any)
		exists = func(key string) bool {
			for k, _ := range store {
				if k == key {
					return true
				}
			}

			return false
		}
		set = func(key string, value any) bool {
			if !storing_errors.CheckIfBadKey(key) {
				return false
			}
			if key == "" {
				storing_errors.NoKeyError("set", key)
				return false
			}

			store[key] = value
			return true
		}
	)

	var initiation *IMap = &IMap{
		All: func() []IMapAll {
			if len(store) < 1 {
				storing_errors.NoKeysError("all", 2)
				return []IMapAll{}
			}

			var res []IMapAll
			for k, v := range store {
				res = append(res, IMapAll{
					Key:   k,
					Value: v,
				})
			}

			return res
		},
		Clear: func() bool {
			if len(store) < 1 {
				storing_errors.NoKeysError("clear", 2)
				return false
			}

			for k, _ := range store {
				delete(store, k)
			}
			return true
		},
		Delete: func(key string) bool {
			if exists(key) {
				delete(store, key)
				return true
			}

			storing_errors.NoKeyError("delete", key)
			return false
		},
		Exists: exists,
		Fetch: func(key string) any {
			_, ok := store[key]
			if ok {
				res := store[key]

				return res
			}

			storing_errors.KeyNotFoundError(key, "get")
			return nil
		},
		Get: func(key string) any {
			for k, v := range store {
				if k == key {
					return v
				}
			}

			storing_errors.KeyNotFoundError(key, "get")
			return nil
		},
		Has: func(key string) bool {
			return exists(key)
		},
		Iterator: func() map[string]any {
			return store
		},
		Keys: func() []string {
			if len(store) < 1 {
				storing_errors.NoKeysError("keys", 2)
				return []string{}
			}

			var res []string
			for k, _ := range store {
				res = append(res, k)
			}

			return res
		},
		Set: set,
		Size: func() int {
			return len(store)
		},
		Update: func(key string, value any) bool {
			if !exists(key) {
				storing_errors.NoKeyError("update", key)

				return false
			}

			return set(key, value)
		},
		Values: func() []any {
			var res []any

			for _, v := range store {
				res = append(res, v)
			}

			return res
		},
	}

	return initiation
}
