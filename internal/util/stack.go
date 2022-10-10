package util

// Stack is an implementation of a LIFO collection
type Stack[T any] []T

// NewStack returns a pointer to a LIFO Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a new element to the top of stack, modifying it.
func (s *Stack[T]) Push(item T) {
	*s = append([]T{item}, *s...)
}

// Pop pulls the most recently pushed element value from the stack and updates the Stack to
// remove that element. Pop returns the element it removes from the stack.
func (s *Stack[T]) Pop() (item T) {
	item, *s = (*s)[0], (*s)[1:]
	return item
}

func (s Stack[T]) Slice() []T {
	return s
}
