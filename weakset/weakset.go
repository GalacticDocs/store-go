package weakset

// Initialises a new WeakSet instance.
func New[T comparable]() *WeakSet[T] {
	return &WeakSet[T]{
		data: make(map[T]*weakValue[T]),
	}
}

func (ws *WeakSet[T]) Add(value T) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	v, ok := ws.data[value]
	if !ok {
		
	}
}