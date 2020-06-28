package stack

// Stack data structure
type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Values() []interface{}
	Length() int
	Empty() bool
	Clear()
}

type stack struct {
	items []interface{}
	size  int
}

// New return array based stack.
// Array based stack good for read operation
func New() Stack {
	return &stack{}
}

// Push item into stack
func (s *stack) Push(item interface{}) {
	s.items = append(s.items, item)
	s.size++
}

// Pop item from stack
func (s *stack) Pop() interface{} {
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
func (s *stack) Peek() interface{} {
	if s.Empty() {
		return nil
	}

	return s.items[s.Length()-1]
}

// Length return lenght of stack
func (s *stack) Length() int {
	return s.size
}

// Clear items in stack
func (s *stack) Clear() {
	s.items = nil
	s.size = 0
}

// Empty return whether stack is empty or not
func (s *stack) Empty() bool {
	return s.Length() == 0
}

// Values return all stack values
func (s *stack) Values() []interface{} {
	return s.items
}
