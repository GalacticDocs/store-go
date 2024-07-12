package weakset

import (
	"fmt"
	"runtime"

	storing_errors "github.com/GalacticDocs/store-go/errors"
)

// Initialises a new WeakSet instance.
func New[T comparable]() *WeakSet[T] {
	return &WeakSet[T]{
		data: make(map[T]*weakValue[T]),
	}
}

// Adds a reference to the WeakSet.
func (ws *WeakSet[T]) Add(value T) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	v, ok := ws.data[value]
	if !ok {
		v = &weakValue[T]{
			refs: make([]*T, 0),
		}
		ws.data[value] = v
	}

	ref := &value
	v.refs = append(v.refs, &*ref)
}

// Checks if a certain reference exists in the current WeakSet.
func (ws *WeakSet[T]) Has(value T) bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	v, ok := ws.data[value]
	if !ok {
		return false
	}

	for _, ref := range v.refs {
		if *ref == value {
			return true
		}
	}

	return false
}

// Deletes a singular reference from the WeakSet.
func (ws *WeakSet[T]) Delete(value T) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if !ws.Has(value) {
		storing_errors.NoKeyError("delete", fmt.Sprint(value))
		return
	}

	delete(ws.data, value)
}

// Deletes all the references from the WeakSet.
func (ws *WeakSet[T]) DeleteWeakRefs() {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for key, v := range ws.data {
		var liveRefs []*T
		for _, ref := range v.refs {
			if ref != nil {
				liveRefs = append(liveRefs, ref)
			}
		}

		v.refs = liveRefs
		if len(v.refs) == 0 {
			delete(ws.data, key)
		}
	}

	runtime.GC()
}