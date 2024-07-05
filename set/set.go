package set

import (
	"fmt"

	storing_errors "github.com/GalacticDocs/store-go/errors"
)

// Initialises a new Cache instance.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// Adds a new value to the Set.
func (s *Set[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[value] = struct{}{}
}

// Checks if a certain value exists in the Set.
func (s *Set[T]) Has(value T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[value]
	return ok
}

// Removes the provided value from the Set. Throws an error if the key doesn't exist.
func (s *Set[T]) Delete(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[value]
	if !ok {
		storing_errors.NoKeyError("delete", fmt.Sprint(value))

		return
	}

	delete(s.data, value)
}

// Returns the size of the Set.
func (s *Set[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[T]struct{})
}

// Returns the values of the current Set.
func (s *Set[T]) Values() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	values := make([]T, len(s.data))
	i := 0
	for key := range s.data {
		values[i] = key
		i++
	}

	return values
}
