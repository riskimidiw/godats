package queue

// Queue data structure
type Queue interface {
	Enqueue(...interface{})
	Dequeue() interface{}
	Peek() interface{}
	Values() []interface{}
	Length() int
	Empty() bool
	Clear()
}
