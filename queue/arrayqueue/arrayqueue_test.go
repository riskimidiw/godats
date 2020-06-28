package arrayqueue

import (
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	tests := []struct {
		Name           string
		Items          []interface{}
		ExpectedLength int
		ExpectedPeek   interface{}
	}{
		{
			Name:           "SingleData",
			Items:          []interface{}{1},
			ExpectedLength: 1,
			ExpectedPeek:   1,
		},
		{
			Name:           "MultipleData",
			Items:          []interface{}{1, 2, 4},
			ExpectedLength: 3,
			ExpectedPeek:   1,
		},
		{
			Name:           "MultipleDataType",
			Items:          []interface{}{3, 2.4, "test"},
			ExpectedLength: 3,
			ExpectedPeek:   3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			stack.Enqueue(test.Items...)

			if stack.Length() != test.ExpectedLength {
				t.Errorf("expected length: %v, got: %v", test.ExpectedLength, stack.Length())
			}
			if stack.Peek() != test.ExpectedPeek {
				t.Errorf("expected peek: %v, got: %v", test.ExpectedPeek, stack.Peek())
			}
		})
	}
}

func BenchmarkQueue_Enqueue_100K(b *testing.B) {
	size := 100000
	stack := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			stack.Enqueue(i)
		}
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		Name           string
		Items          []interface{}
		ExpectedResult interface{}
		ExpectedLength int
		ExpectedPeek   interface{}
	}{
		{
			Name:           "EmptyData",
			ExpectedResult: nil,
			ExpectedLength: 0,
			ExpectedPeek:   nil,
		},
		{
			Name:           "SingleData",
			Items:          []interface{}{1},
			ExpectedResult: 1,
			ExpectedLength: 0,
			ExpectedPeek:   nil,
		},
		{
			Name:           "MultipleData",
			Items:          []interface{}{1, 2, 4},
			ExpectedResult: 1,
			ExpectedLength: 2,
			ExpectedPeek:   2,
		},
		{
			Name:           "MultipleDataType",
			Items:          []interface{}{1, 2.4, "test"},
			ExpectedResult: 1,
			ExpectedLength: 2,
			ExpectedPeek:   2.4,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			stack.Enqueue(test.Items...)

			item := stack.Dequeue()
			if item != test.ExpectedResult {
				t.Errorf("expected result: %v, got: %v", test.ExpectedResult, item)
			}
			if stack.Length() != test.ExpectedLength {
				t.Errorf("expected length: %v, got: %v", test.ExpectedLength, stack.Length())
			}
			if stack.Peek() != test.ExpectedPeek {
				t.Errorf("expected peek: %v, got: %v", test.ExpectedPeek, stack.Peek())
			}
		})
	}
}

func BenchmarkQueue_Dequeue_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			stack.Dequeue()
		}
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		Name         string
		Items        []interface{}
		ExpectedPeek interface{}
	}{
		{
			Name:         "EmptyData",
			ExpectedPeek: nil,
		},
		{
			Name:         "SingleData",
			Items:        []interface{}{1},
			ExpectedPeek: 1,
		},
		{
			Name:         "MultipleData",
			Items:        []interface{}{1, 2, 4},
			ExpectedPeek: 1,
		},
		{
			Name:         "MultipleDataType",
			Items:        []interface{}{3, 2.4, "test"},
			ExpectedPeek: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			stack.Enqueue(test.Items...)

			item := stack.Peek()
			if item != test.ExpectedPeek {
				t.Errorf("expected peek: %v, got: %v", test.ExpectedPeek, item)
			}
		})
	}
}

func BenchmarkQueue_Peek_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Enqueue(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

func TestQueue_Lenght(t *testing.T) {
	tests := []struct {
		Name           string
		Items          []interface{}
		ExpectedLength interface{}
	}{
		{
			Name:           "EmptyData",
			ExpectedLength: 0,
		},
		{
			Name:           "SingleData",
			Items:          []interface{}{1},
			ExpectedLength: 1,
		},
		{
			Name:           "MultipleData",
			Items:          []interface{}{1, 2, 4},
			ExpectedLength: 3,
		},
		{
			Name:           "MultipleDataType",
			Items:          []interface{}{1, 2.4, "test"},
			ExpectedLength: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			stack.Enqueue(test.Items...)

			result := stack.Length()
			if result != test.ExpectedLength {
				t.Errorf("expected length: %v, got: %v", test.ExpectedLength, result)
			}
		})
	}
}

func BenchmarkQueue_Length_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Length()
	}
}

func TestQueue_Empty(t *testing.T) {
	tests := []struct {
		Name          string
		Items         []interface{}
		ExpectedEmpty bool
	}{
		{
			Name:          "EmptyData",
			ExpectedEmpty: true,
		},
		{
			Name:          "SingleData",
			Items:         []interface{}{1},
			ExpectedEmpty: false,
		},
		{
			Name:          "MultipleData",
			Items:         []interface{}{1, 2, 4},
			ExpectedEmpty: false,
		},
		{
			Name:          "MultipleDataType",
			Items:         []interface{}{1, 2.4, "test"},
			ExpectedEmpty: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			stack.Enqueue(test.Items...)

			result := stack.Empty()
			if result != test.ExpectedEmpty {
				t.Errorf("expected result: %v, got: %v", test.ExpectedEmpty, result)
			}
		})
	}
}

func BenchmarkQueue_Empty_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Empty()
	}
}

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		Name          string
		Items         []interface{}
		ExpectedEmpty bool
	}{
		{
			Name: "EmptyData",
		},
		{
			Name:  "SingleData",
			Items: []interface{}{1},
		},
		{
			Name:  "MultipleData",
			Items: []interface{}{1, 2, 4},
		},
		{
			Name:  "MultipleDataType",
			Items: []interface{}{1, 2.4, "test"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			stack.Enqueue(test.Items...)

			stack.Clear()
			empty := stack.Empty()
			if !empty {
				t.Errorf("expected empty: %v, got: %v", true, empty)
			}
			length := stack.Length()
			if length != 0 {
				t.Errorf("expected length: %v, got: %v", 0, length)
			}
		})
	}
}

func BenchmarkQueue_Clear_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Clear()
	}
}

func TestQueue_Values(t *testing.T) {
	tests := []struct {
		Name  string
		Items []interface{}
	}{
		{
			Name: "EmptyData",
		},
		{
			Name:  "SingleData",
			Items: []interface{}{1},
		},
		{
			Name:  "MultipleData",
			Items: []interface{}{1, 2, 4},
		},
		{
			Name:  "MultipleDataType",
			Items: []interface{}{1, 2.4, "test"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			stack.Enqueue(test.Items...)

			item := stack.Values()
			if !reflect.DeepEqual(item, test.Items) {
				t.Errorf("expected result: %v, got: %v", test.Items, item)
			}
		})
	}
}

func BenchmarkQueue_Values_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Values()
	}
}
