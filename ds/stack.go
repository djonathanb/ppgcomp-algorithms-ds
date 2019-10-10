package ds

import "errors"

// Stack LIFO data structure
type Stack struct {
	vet []int
	tos int
}

// NewStack returns a new Stack given a determined size
func NewStack(size int) Stack {
	return Stack{
		vet: make([]int, size),
		tos: -1,
	}
}

// Count return the number of items of the stack
func (s *Stack) Count() int {
	return s.tos + 1
}

// Push insert data at the top of the Stack
func (s *Stack) Push(data int) {
	s.tos = s.tos + 1
	s.vet[s.tos] = data
}

// Pop remove and return the item on the top of the Stack
func (s *Stack) Pop() (int, error) {
	if s.tos >= 0 {
		temp := s.vet[s.tos]
		s.tos = s.tos - 1
		return temp, nil
	}
	return -1, errors.New("empty stack")
}

// IsEmpty return true if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Count() == 0
}
