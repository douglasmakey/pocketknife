package datatype

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zv T
		return zv, false
	}

	index := len(*s) - 1
	v := (*s)[index]
	*s = (*s)[:index]
	return v, true
}
