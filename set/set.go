package set

func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[value] = struct{}{}
}