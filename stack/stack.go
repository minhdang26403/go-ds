package stack

import "fmt"

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Top() (value T, err error) {
	if stack.IsEmpty() {
		return value, fmt.Errorf("The stack is empty")
	}
	value = stack.data[len(stack.data) - 1]
	return value, err
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack[T]) Size() int {
	return len(stack.data)
}

func (stack *Stack[T]) Push(value T) {
	stack.data = append(stack.data, value)
}

func (stack *Stack[T]) Pop() (value T, err error) {
	value, err = stack.Top()
	if err != nil {
		return value, err
	}
	stack.data = stack.data[:len(stack.data) - 1]
	return value, err
}