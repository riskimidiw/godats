package arrayqueue

import (
	"github.com/riskimidiw/godats/queue"
)

type arrayQueue struct {
	items []interface{}
	size  int
}

// New return array based queue.
// Array based queue good for read operation
func New() queue.Queue {
	return &arrayQueue{}
}

// Enqueue item into queue
func (s *arrayQueue) Enqueue(items ...interface{}) {
	for _, item := range items {
		s.items = append(s.items, item)
		s.size++
	}
}

// Dequeue item from queue
func (s *arrayQueue) Dequeue() interface{} {
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
func (s *arrayQueue) Peek() interface{} {
	if s.Empty() {
		return nil
	}

	return s.items[0]
}

// Length return lenght of queue
func (s *arrayQueue) Length() int {
	return s.size
}

// Clear items in queue
func (s *arrayQueue) Clear() {
	s.items = nil
	s.size = 0
}

// Empty return whether queue is empty or not
func (s *arrayQueue) Empty() bool {
	return s.Length() == 0
}

// Values return all queue values from earliest queued item
func (s *arrayQueue) Values() []interface{} {
	return s.items
}
