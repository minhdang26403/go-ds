package queue

import "github.com/minhdang26403/algo-ds/linkedlist"

type Queue[T any] struct {
	data linkedlist.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: *linkedlist.NewLinkedList[T]()}
}

func (queue *Queue[T]) Size() int {
	return queue.data.Size()
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *Queue[T]) Front() (T, error) {
	return queue.data.Front()
}

func (queue *Queue[T]) Back() (T, error) {
	return queue.data.Back()
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.data.PushBack(value)
}

func (queue *Queue[T]) Dequeue() (T, error) {
	value, err := queue.data.Front()
	queue.data.PopFront()
	return value, err
}