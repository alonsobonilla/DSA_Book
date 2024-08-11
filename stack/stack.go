package stack

type stack[T any] struct {
	topOfStack int
	arr        []T
}

func NewStack[T any](capacity int) *stack[T] {
	return &stack[T]{
		topOfStack: -1,
		arr:        make([]T, capacity),
	}
}

func (s *stack[T]) Pop() T {
	element := s.arr[s.topOfStack]
	s.topOfStack--
	return element
}

func (s *stack[T]) Top() (T, bool) {
	if s.topOfStack < 0 {
		var zeroValue T
		return zeroValue, false
	}
	return s.arr[s.topOfStack], true
}

func (s *stack[T]) Push(element T) {
	s.topOfStack++
	s.arr[s.topOfStack] = element
}
