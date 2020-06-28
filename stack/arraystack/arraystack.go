package arraystack

import (
	"github.com/riskimidiw/godats/stack"
)

type arrayStack struct {
	items []interface{}
	size  int
}

// New return array based stack.
// Array based stack good for read operation
func New() stack.Stack {
	return &arrayStack{}
}

// Push item into stack
func (s *arrayStack) Push(items ...interface{}) {
	for _, item := range items {
		s.items = append(s.items, item)
		s.size++
	}
}

// Pop item from stack
func (s *arrayStack) Pop() interface{} {
	if s.Empty() {
		return nil
	}

	item := s.items[s.Length()-1]
	s.size--
	s.items = s.items[:s.Length()]
	return item
}

// Peek return the first element of the stack
// or the element present at the top of the stack
func (s *arrayStack) Peek() interface{} {
	if s.Empty() {
		return nil
	}

	return s.items[s.Length()-1]
}

// Length return lenght of stack
func (s *arrayStack) Length() int {
	return s.size
}

// Clear items in stack
func (s *arrayStack) Clear() {
	s.items = nil
	s.size = 0
}

// Empty return whether stack is empty or not
func (s *arrayStack) Empty() bool {
	return s.Length() == 0
}

// Values return all stack values from earliest pushed item
func (s *arrayStack) Values() []interface{} {
	return s.items
}
