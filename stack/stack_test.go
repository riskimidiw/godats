package stack

import (
	"reflect"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
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
			ExpectedPeek:   4,
		},
		{
			Name:           "MultipleDataType",
			Items:          []interface{}{1, 2.4, "test"},
			ExpectedLength: 3,
			ExpectedPeek:   "test",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			for _, item := range test.Items {
				stack.Push(item)
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

func BenchmarkArrayStack_Push_100K(b *testing.B) {
	size := 100000
	stack := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			stack.Push(i)
		}
	}
}

func TestArrayStack_Pop(t *testing.T) {
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
			ExpectedResult: 4,
			ExpectedLength: 2,
			ExpectedPeek:   2,
		},
		{
			Name:           "MultipleDataType",
			Items:          []interface{}{1, 2.4, "test"},
			ExpectedResult: "test",
			ExpectedLength: 2,
			ExpectedPeek:   2.4,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			for _, item := range test.Items {
				stack.Push(item)
			}

			item := stack.Pop()
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

func BenchmarkArrayStack_Pop_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			stack.Pop()
		}
	}
}

func TestArrayStack_Peek(t *testing.T) {
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
			ExpectedPeek: 4,
		},
		{
			Name:         "MultipleDataType",
			Items:        []interface{}{1, 2.4, "test"},
			ExpectedPeek: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := New()
			for _, item := range test.Items {
				stack.Push(item)
			}

			item := stack.Peek()
			if item != test.ExpectedPeek {
				t.Errorf("expected peek: %v, got: %v", test.ExpectedPeek, item)
			}
		})
	}
}

func BenchmarkArrayStack_Peek_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Push(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

func TestArrayStack_Lenght(t *testing.T) {
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
			for _, item := range test.Items {
				stack.Push(item)
			}

			result := stack.Length()
			if result != test.ExpectedLength {
				t.Errorf("expected length: %v, got: %v", test.ExpectedLength, result)
			}
		})
	}
}

func BenchmarkArrayStack_Lengh_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Length()
	}
}

func TestArrayStack_Empty(t *testing.T) {
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
			for _, item := range test.Items {
				stack.Push(item)
			}

			result := stack.Empty()
			if result != test.ExpectedEmpty {
				t.Errorf("expected result: %v, got: %v", test.ExpectedEmpty, result)
			}
		})
	}
}

func BenchmarkArrayStack_Empty_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Empty()
	}
}

func TestArrayStack_Clear(t *testing.T) {
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
			for _, item := range test.Items {
				stack.Push(item)
			}

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

func BenchmarkArrayStack_Clear_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Clear()
	}
}

func TestArrayStack_Values(t *testing.T) {
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
			for _, item := range test.Items {
				stack.Push(item)
			}

			item := stack.Values()
			if !reflect.DeepEqual(item, test.Items) {
				t.Errorf("expected result: %v, got: %v", test.Items, item)
			}
		})
	}
}

func BenchmarkArrayStack_Values_100K(b *testing.B) {
	size := 100000
	stack := New()
	for i := 0; i < size; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Values()
	}
}
