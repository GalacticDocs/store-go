package collection_store

import (
	"sync"

	storing_errors "github.com/GalacticDocs/store-go/errors"
)

// New creates a new Collection instance and returns it.
//
// Returns a new Collection instance.
func New() *ICollection {
	return &ICollection{
		store: sync.Map{},
	}
}

// Clear removes all elements from the Collection.
//
// Returns a boolean indicating whether the Collection has been cleared or not.
func (c *ICollection) Clear() bool {
	c.store.Range(func(key, value any) bool {
		c.store.Delete(key)
		return true
	})

	return true
}

// Delete removes any value associated to the key from the Collection.
//
// The Delete function takes 1 parameter:
//   - key: a string representing the key to search for, which is associated to a string type.
//
// The first parameter represents the key to search for, which is associated to a string type.
//
// Returns a boolean indicating whether the key-value pair has been deleted or not.
func (c *ICollection) Delete(key string) bool {
	_, loaded := c.store.LoadAndDelete(key)

	if !loaded {
		storing_errors.NoKeyError("delete", key)
		return false
	}

	return true
}

// Exists returns a Boolean asserting whether a value has been associated to the key in the Collection or not.
//
// The Exists function takes 1 parameter:
//   - key: a string representing the key to search for, which is associated to a string type.
//
// The first parameter represents the key to search for, which is associated to a string type.
//
// Returns a Boolean whether or not the key exists in the Collection.
func (c *ICollection) Exists(key string) bool {
	_, loaded := c.store.Load(key)

	return loaded
}

// Fetch returns the value associated to the key, or nil if there is none.
//
// The Fetch function takes 1 parameter:
//   - key: a string representing the key to search for, which is associated to a string type.
//
// The first parameter represents the key to search for, which is associated to a string type.
//
// Returns the value associated to the key, or nil if there is none.
func (c *ICollection) Fetch(key string) any {
	data, _ := c.store.Load(key)

	return data
}

// Filter applies the given function to each key-value pair in the Collection and returns a new Collection containing
// only the key-value pairs for which the function returns true.
//
// The Filter function takes 1 parameter:
//   - fn: a function to be applied to each key-value pair in the Collection.
//
// The first parameter (fn) represents a function to be applied to each key-value pair in the Collection.
// The function returns a boolean indicating whether the key-value pair should be
// included in the resulting ICollection.
//
// The fn function takes three parameters:
//   - key: a string representing the key of the key-value pair.
//   - value: the value associated with the key.
//   - collection: a pointer to the Collection being filtered.
//
// The function returns a pointer to the new ICollection containing the filtered key-value pairs.
func (c *ICollection) Filter(fn IFilterFunc) *ICollection {
	var result = New()

	c.store.Range(func(key, value any) bool {
		if fn(key.(string), value, c) {
			result.store.Store(key, value)
		}

		return true
	})

	c = result
	return c
}

// Find searches for a value in the Collection based on the provided function.
//
// The Find function takes 1 parameter:
//   - fn: a function which is used to determine whether a value should be considered a match or not.
//
// The first parameter represents a function to be applied to each key-value pair in the Collection.
// The function returns a boolean indicating whether the key-value pair should be considered a match.
//
// The fn function takes three parameters:
//   - key: a string representing the key of the key-value pair.
//   - value: the value associated with the key.
//   - collection: a pointer to the Collection being searched.
//
// Returns the value associated with the key that matches the provided function, or nil if no match is found.
func (c *ICollection) Find(fn IFindFunc) any {
	var result any
	c.store.Range(func(key, value any) bool {
		if fn(key.(string), value, c) {
			result = value
			return true
		} else {
			result = nil
			return false
		}
	})

	return result
}

// First fetches the first item in the Collection.
//
// Returns the first item in the Collection.
func (c *ICollection) First() any {
	var result any

	c.store.Range(func(key, value any) bool {
		ind := c.GetIndex(key.(string))

		if ind == 0 {
			result = value
			return true
		} else {
			result = nil
			return true
		}
	})

	return result
}

// Get returns the value associated to the key, or nil if there is none.
//
// The Get function takes 1 parameter:
//   - key: a string representing the key to search for, which is associated to a string type.
//
// The first parameter represents the key to search for, which is associated to a string type.
//
// Returns the value associated to the key, or nil if there is none.
func (c *ICollection) Get(key string) any {
	return c.Fetch(key)
}

// GetIndex returns the index of the given key in the collection.
//
// The GetIndex function takes 1 parameter:
//   - key: a string representing the key to search for, which is associated to a string type.
//
// The first parameter represents the key to search for, which is associated to a string type.
//
// Returns the index of the given key in the collection.
func (c *ICollection) GetIndex(key string) int {
	for i, k := range c.ToKeyArray() {
		if k == key {
			return i
		}
	}

	return -1
}

// Has returns a Boolean asserting whether a value has been associated to the key in the Collection or not.
//
// The Has function takes 1 parameter:
//   - key: a string representing the key to search for, which is associated to a string type.
//
// The first parameter represents the key to search for, which is associated to a string type.
//
// Returns a Boolean whether or not the key exists in the Collection.
func (c *ICollection) Has(key string) bool {
	_, loaded := c.store.Load(key)
	return loaded
}

// Iterator returns a map containing the elements of the ICollection.
//
// Returns a regular map which can be used to iterate over.
func (c *ICollection) Iterator() map[string]any {
	var result = make(map[string]any)
	c.store.Range(func(key, value any) bool {
		result[key.(string)] = value
		return true
	})

	return result
}

// Last fetches the last item in the Collection.
//
// Returns the last item in the Collection.
func (c *ICollection) Last() any {
	var result any
	c.store.Range(func(key, value any) bool {
		ind := c.GetIndex(key.(string))

		if ind == len(c.ToKeyArray())-1 {
			result = value
			return true
		} else {
			result = nil
			return true
		}
	})

	return result
}

// Map applies a function to each key-value pair in the Collection and returns an array of the results.
//
// The Map function takes 1 parameter:
//   - fn: a function that maps the key-value pair.
//
// The first parameter represents a function that maps the key-value pair.
//
// The fn function takes three parameters:
//   - key: a string representing the key of the key-value pair.
//   - value: the value associated with the key.
//   - collection: a pointer to the Collection being mapped.
//
// Returns an array of the results.
func (c *ICollection) Map(fn IMapTransformer) []any {
	var result []any
	var arrayIteration int

	c.store.Range(func(key, value any) bool {
		result = append(result, fn(value, key.(string), arrayIteration, c))
		arrayIteration++
		return true
	})

	return result
}

// Reduce applies a function against an accumulator and each value of the Collection (from left-to-right) to reduce it to a single value.
//
// The Reduce function takes 2 parameters:
//   - fn: a function which is used to reduce the Collection.
//   - initialAccumulator: an initial accumulator value.
//
// The first parameter represents a function to be applied to each key-value pair in the Collection.
// The function returns an accumulator value.
//
// The second parameter represents the initial accumulator value.
//
// The fn function takes three parameters:
//   - accumulator: an accumulator value.
//   - value: the value associated with the key.
//   - key: a string representing the key of the key-value pair.
//   - collection: a pointer to the Collection being reduced.
//
// Returns the accumulator value.
func (c *ICollection) Reduce(fn IReduceTransformer, initialAccumulator any) any {
	var accumulator any
	var isFirstItem = true

	if initialAccumulator != nil {
		accumulator = initialAccumulator
		c.store.Range(func(key, value any) bool {
			accumulator = fn(accumulator, value, key.(string), c)
			return true
		})
	} else {
		c.store.Range(func(key, value any) bool {
			if isFirstItem {
				accumulator = value
				isFirstItem = false
			} else {
				accumulator = fn(accumulator, value, key.(string), c)
			}

			return true
		})
	}

	return accumulator
}

// Set sets the key-value pair in the Collection.
//
// The Set function takes 2 parameters:
//   - key: a string representing the key of the key-value pair.
//   - value: the value associated with the key.
//
// The first parameter represents a string representing the key of the key-value pair.
//
// The second parameter represents the value associated with the key.
//
// Returns a pointer to the Collection.
func (c *ICollection) Set(key string, value any) *ICollection {
	c.store.Store(key, value)

	return c
}

// Size returns the amount of keys/values in the Collection.
//
// Returns the amount of keys in the Collection.
func (c *ICollection) Size() int {
	var size int = 0
	c.store.Range(func(key, value any) bool {
		size++
		return true
	})

	return size
}

// Some applies a function checks if at least one element in the Collection satisfies the given function.
//
// The Some function takes 1 parameter:
//   - fn: a function that checks if an element satisfies the provided filter.
//
// The first parameter represents a function that checks if an element satisfies the provided filter.
//
// The fn function takes three parameters:
//   - key: a string representing the key of the key-value pair.
//   - value: the value associated with the key.
//   - collection: a pointer to the Collection being filtered.
//
// Returns a Boolean asserting whether at least one element in the Collection satisfies the given function.
func (c *ICollection) Some(fn ISomeFunc) bool {
	var result bool = false
	
	c.store.Range(func(key, value any) bool {
		result = fn(key.(string), value, c)
		return true
	})

	return result
}

// Sorts the items in the collection
func (c *ICollection) Sort(fn func(a any, b any)) *ICollection {
	//
}

// Sweep removes all elements in the Collection that match the given function.
//
// The Sweep function takes 1 parameter:
//   - fn: a function that removes an element based on the provided filter.
//
// The first parameter represents a function that removes an element based on the provided filter.
//
// The fn function takes three parameters:
//   - key: a string representing the key of the key-value pair.
//   - value: the value associated with the key.
//   - collection: a pointer to the Collection being filtered.
//
// Returns the amount of elements removed from the Collection.
func (c *ICollection) Sweep(fn ISweepFunc) int {
	previousSize := c.Size()

	c.store.Range(func(key, value any) bool {
		if fn(key.(string), value, c) {
			c.store.Delete(key)
		}

		return true
	})

	return previousSize - c.Size()
}

// ToArray returns all the key-value pairs in the Collection as an array of the struct ICollectionAll.
//
// Returns an array indicating all the key-value pairs in the Collection.
func (c *ICollection) ToArray() []ICollectionAll {
	var result []ICollectionAll
	c.store.Range(func(key, value any) bool {
		result = append(result, ICollectionAll{Key: key.(string), Value: value})
		return true
	})

	return result
}

// ToKeyArray returns an array of strings containing all the keys in the Collection.
//
// Returns an array of strings containing all the keys in the Collection.
func (c *ICollection) ToKeyArray() []string {
	var result []string
	c.store.Range(func(key, value any) bool {
		result = append(result, key.(string))
		return true
	})

	return result
}
