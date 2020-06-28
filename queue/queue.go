package queue

// Queue data structure
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Peek() interface{}
	Values() []interface{}
	Length() int
	Empty() bool
	Clear()
}

type queue struct {
	items []interface{}
	size  int
}

// New return array based queue.
// Array based queue good for read operation
func New() Queue {
	return &queue{}
}

// Enqueue item into queue
func (s *queue) Enqueue(item interface{}) {
	s.items = append(s.items, item)
	s.size++
}

// Dequeue item from queue
func (s *queue) Dequeue() interface{} {
	if s.Empty() {
		return nil
	}

	item := s.items[0]
	s.size--
	s.items = s.items[1:]
	return item
}

// Peek return the first element of the queue
// or the element present at the head of the queue
func (s *queue) Peek() interface{} {
	if s.Empty() {
		return nil
	}

	return s.items[0]
}

// Length return lenght of queue
func (s *queue) Length() int {
	return s.size
}

// Clear items in queue
func (s *queue) Clear() {
	s.items = nil
	s.size = 0
}

// Empty return whether queue is empty or not
func (s *queue) Empty() bool {
	return s.Length() == 0
}

// Values return all queue values from earliest queued item
func (s *queue) Values() []interface{} {
	return s.items
}
