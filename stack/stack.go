package stack

type Stack[T any] struct {
	topOfStack int
	arr        []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		topOfStack: -1,
		arr:        make([]T, capacity),
	}
}

func (s *Stack[T]) Pop() T {
	element := s.arr[s.topOfStack]
	s.topOfStack--
	return element
}

func (s *Stack[T]) Top() T {
	if s.topOfStack < 0 {
		var zeroValue T
		return zeroValue
	}
	return s.arr[s.topOfStack]
}

func (s *Stack[T]) Push(element T) {
	s.topOfStack++
	s.arr[s.topOfStack] = element
}

func (s Stack[T]) Empty() bool {
	return s.topOfStack < 0
}
