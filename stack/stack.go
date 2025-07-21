package stack

// Stack represents a LIFO (Last In, First Out) data structure for storing elements.
type Stack struct {
	data []any
}

// NewStack creates and returns a new instance of Stack initialized with an empty data slice.
func NewStack() *Stack {
	return &Stack{data: []any{}}
}

// Push adds a new element to the top of the stack.
func (s *Stack) Push(value any) {
	s.data = append(s.data, value)
}

// Pop removes the top element from the stack and returns it. Returns nil if the stack is empty.
func (s *Stack) Pop() any {
	if len(s.data) == 0 {
		return nil
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

// Top returns the top element of the stack without removing it. Returns nil if the stack is empty.
func (s *Stack) Top() any {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}

// Dump returns a slice containing all elements in the stack in their current order.
func (s *Stack) Dump() []any {
	return s.data
}

// Reset clears all elements from the stack, leaving it empty.
func (s *Stack) Reset() {
	s.data = s.data[:0]
}

// Peek returns the top element of the stack without removing it. Returns nil if the stack is empty.
func (s *Stack) Peek() any {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

// Len returns the number of elements currently stored in the stack.
func (s *Stack) Len() int {
	return len(s.data)
}

// PeekNFromTop retrieves the nth element from the top of the stack without removing it. Returns nil if index is out of bounds.
func (s *Stack) PeekNFromTop(n int) any {
	idx := len(s.data) - 1 - n
	if idx < 0 || idx >= len(s.data) {
		return nil
	}
	return s.data[idx]
}
