package stack

// Stack data structure
type Stack interface {
	Push(...interface{})
	Pop() interface{}
	Peek() interface{}
	Values() []interface{}
	Length() int
	Empty() bool
	Clear()
}
