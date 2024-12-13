package map_store

import (
	"sort"
	"sync"

	storing_errors "github.com/GalacticDocs/store-go/errors"
)

func Map() *IMap {
	return &IMap{
		store: sync.Map{},
	}
}

// Returns all the existing keys in the current Map.
func (m *IMap) All() []IMapAll {
	var result []IMapAll
	m.store.Range(func(key, value any) bool {
		result = append(result, IMapAll{Key: key.(string), Value: value})
		return true
	})

	return result
}

// The Clear() method removes all elements from the Map. Returns whether the map has been cleared or not.
func (m *IMap) Clear() bool {
	m.store.Range(func(key, value any) bool {
		m.store.Delete(key)
		return true
	})

	return true
}

// Removes any value associated to the key. Map().has(key) will return false afterwards.
func (m *IMap) Delete(key string) bool {
	_, loaded := m.store.LoadAndDelete(key)

	if !loaded {
		storing_errors.NoKeyError("delete", key)
		return false
	}

	return true
}

// Returns a Boolean asserting whether a value has been associated to the key in the Map or not.
func (m *IMap) Exists(key string) bool {
	_, loaded := m.store.Load(key)
	return loaded
}

// Returns the value associated to the key, or nil if there is none.
func (m *IMap) Fetch(key string) any {
	data, _ := m.store.Load(key)

	return data
}

// First searches the first value in the map.
// Returns the first value in the map.
func (m *IMap) First() any {
	var result any

	m.store.Range(func(key, value any) bool {
		result = value
		return false
	})

	return result
}

// Get retrieves the value associated with the given key from the Map.
// It returns the value if found, or nil if no value is associated with the key.
func (m *IMap) Get(key string) any {
	return m.Fetch(key)
}

// Returns a Boolean asserting whether a value has been associated to the key in the Map or not.
func (m *IMap) Has(key string) bool {
	_, loaded := m.store.Load(key)
	return loaded
}

// The iterator which is used to iterate through the map.
func (m *IMap) Iterator() map[string]any {
	var result = make(map[string]any)
	m.store.Range(func(key, value any) bool {
		result[key.(string)] = value
		return true
	})

	return result
}

// Returns all the keys as an array.
func (m *IMap) Keys() []string {
	var keys []string
	m.store.Range(func(key, value any) bool {
		keys = append(keys, key.(string))
		return true
	})

	return keys
}

// Sets the value for the key in the Map. Returns whether the key value has been set.
func (m *IMap) Set(key string, value any) bool {
	if m.Has(key) {
		_, loaded := m.store.LoadOrStore(key, value)
		return loaded
	} else {
		m.store.Store(key, value)
		return true
	}
}

// Returns the amount of keys in the Map.
func (m *IMap) Size() int {
	var size int
	m.store.Range(func(key, value any) bool {
		size++
		return true
	})

	return size
}

func (m *IMap) Sort(fn func(a any, b any) bool) *IMap {
	// Extract key-value pairs from the sync.Map
    keyValuePairs := make([][2]any, 0)
    m.store.Range(func(key, value any) bool {
        keyValuePairs = append(keyValuePairs, [2]any{key, value})
        return true
    })

    // Sort the key-value pairs based on the comparison function
    sort.Slice(keyValuePairs, func(i, j int) bool {
        firstVal := keyValuePairs[i][1]
		secondVal := keyValuePairs[j][1]

		return fn(firstVal, secondVal)
    })

	// Clear the original map
    m.store.Range(func(key, _ any) bool {
        m.store.Delete(key)
        return true
    })

	// Re-add the sorted key-value pairs to the original map
    for _, pair := range keyValuePairs {
        m.store.Store(pair[0], pair[1])
    }

    return m
}

// Returns all the values as an array.
func (m *IMap) Values() []any {
	var values []any
	m.store.Range(func(key, value any) bool {
		values = append(values, value)
		return true
	})

	return values
}
