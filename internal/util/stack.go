package util

// Stack is an implementation of a LIFO collection
type Stack[T any] []T

// NewStack returns a pointer to a LIFO Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a new element to the top of stack, modifying it.
func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

// Pop pulls the most recently pushed element value from the stack and updates the Stack to
// remove that element. Pop returns the element it removes from the stack.
func (s *Stack[T]) Pop() (item T) {
	if len(*s) == 0 {
		return item
	}
	*s, item = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return item 
}

// Slice returns the underlying slice of the Stack
func (s *Stack[T]) Slice() []T {
	return *s
}

// Clone returns a copy of the Stack
func (s *Stack[T]) Clone() *Stack[T] {
	clone := NewStack[T]()
	for _, item := range Reverse(s.Slice()) {
		clone.Push(item)
	}

	return clone
}
