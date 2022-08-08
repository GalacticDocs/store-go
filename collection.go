package store

import (
	"fmt"

	storing_errors "github.com/GalacticDocs/store-go/errors"
)

type IHasAll struct {
	// Returns if all keys exist.
	AllExist bool
	// Returns the keys that don't exist.
	Keys []string
}

type IHasAny struct {
	HasAny bool
	Keys   []string
}

type ICollection struct {
	// Returns all the existing keys in the Collection.
	All func() []IMapAll
	// Returns the item at a given index, allowing for positive and negative integers. If it fails then it'll return nil.
	At func(idx int) any
	// The Clear() method removes all elements from the Collection. Returns whether the collection has been cleared or not.
	Clear func() bool
	// Removes any value associated to the key. Collection().has(key) will return false afterwards.
	Delete func(key string) bool
	// Returns a Boolean asserting whether a value has been associated to the key in the Collection or not.
	Exists func(key string) bool
	// Returns the value associated to the key, or nil if there is none.
	Fetch func(key string) any
	// Returns the first value from this Collection. If there is none or if it fails then it'll return nil.
	First func() any
	// Returns the first associated key in the Collection. If there is none or if it fails then it'll return an empty string.
	FirstKey func() string
	// Runs the given function in a for-loop/for-each-loop.
	ForEach func(fn func(value any, key string) any)
	// Returns the value associated to the key, or nil if there is none. NOTE: for stricter fetching you may want to consider using Collection().fetch(key)
	Get func(key string) any
	// Returns a Boolean asserting whether a value has been associated to the key in the Collection or not.
	Has func(key string) bool
	// Checks if all of the elements exist in the Collection.
	HasAll func(keys []string) *IHasAll
	// Checks if any of the provided keys exist in the Collection. If one exists it'll return if any exist and the key that exists.
	HasAny func(keys []string) *IHasAny
	// The iterator which is used to iterate through the map.
	Iterator func() map[string]any
	// Joins the values with the given seperator.
	Join func(seperator string) string
	// Joins the keys with the given seperator.
	JoinKeys func(seperator string) string
	// Returns all the keys as an array.
	Keys func() []string
	// Returns the key at a given index, allowing for positive and negative integers. If it fails then it'll return an empty string.
	KeyAt func(idx int) string
	// Returns the last value from this Collection. If there is none or if it fails then it'll return nil.
	Last func() any
	// Returns the last associated key in the Collection. If there is none or if it fails then it'll return an empty string.
	LastKey func() string
	// Sets the value for the key in the Collection. Returns whether the key value has been set.
	Set func(key string, value any) bool
	// Returns the amount of keys in the Collection.
	Size func() int
	// Updates the value for the provided key in the Collection. Returns whether the key value has been changed.
	Update func(key string, value any) bool
	// Returns all the values as an array.
	Values func() []any
}

var store = Map()

// Initiates a new collection instance.
func Collection() *ICollection {
	var (
		initiation *ICollection = &ICollection{
			All:   store.All,
			Clear: store.Clear,
			At: func(idx int) any {
				for _idx, i := range store.All() {
					if _idx == idx {
						return i.Value
					}
				}

				return nil
			},
			Delete: store.Delete, 
			Exists: store.Exists,
			Fetch: store.Fetch,
			First: func() any {
				for idx, i := range store.All() {
					if idx == (1 - 1) {
						return i.Value
					}
				}

				return nil
			},
			FirstKey: func() string {
				for idx, i := range store.All() {
					if idx == (1 - 1) {
						return i.Key
					}
				}

				return ""
			},
			ForEach: func(fn func(value any, key string) any) {
				for _, i := range store.All() {
					fn(i.Value, i.Key)
				}
			},
			Get: store.Get,
			Has: store.Has,
			HasAll: func(keys []string) *IHasAll {
				var (
					existCheck bool
					keyChain   []string
				)

				for _, i := range keys {
					if !store.Exists(i) {
						keyChain = append(keyChain, i)
					}
				}

				if len(keyChain) < 1 {
					existCheck = false
				} else if len(keyChain) > 0 {
					existCheck = true
				}

				return &IHasAll{
					Keys:     keyChain,
					AllExist: existCheck,
				}
			},
			HasAny: func(keys []string) *IHasAny {
				var (
					existCheck bool
					keyChain   []string
				)

				for _, i := range keys {
					if store.Exists(i) {
						keyChain = append(keyChain, i)
					}
				}

				if len(keyChain) < 1 {
					existCheck = false
				} else if len(keyChain) > 0 {
					existCheck = true
				}

				return &IHasAny{
					Keys:   keyChain,
					HasAny: existCheck,
				}
			},
			Iterator: store.Iterator,
			Join: func(seperator string) string {
				if !storing_errors.CheckIfBadJoinSeperator(seperator) {
					return ""
				}

				var str string
				for _, i := range store.All() {
					str = str + fmt.Sprint(i.Value) + seperator
				}

				return str
			},
			JoinKeys: func(seperator string) string {
				if !storing_errors.CheckIfBadJoinSeperator(seperator) {
					return ""
				}

				var str string
				for _, i := range store.All() {
					str = str + fmt.Sprint(i.Value) + seperator
				}

				return str
			},
			Keys: store.Keys,
			KeyAt: func(idx int) string {
				for _idx, i := range store.All() {
					if _idx == idx {
						return i.Key
					}
				}

				return ""
			},
			Last: func() any {
				for idx, i := range store.All() {
					if idx == (len(store.Iterator()) - 1) {
						return i.Value
					}
				}

				return nil
			},
			LastKey: func() string {
				for idx, i := range store.All() {
					if idx == (len(store.Iterator()) - 1) {
						return i.Key
					}
				}

				return ""
			},
			Set:    store.Set,
			Size:   store.Size,
			Update: store.Update,
			Values: store.Values,
		}
	)

	return initiation
}
