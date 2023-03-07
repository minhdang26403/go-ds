package queue

import list "github.com/minhdang26403/algo-ds/linkedlist"

type Queue struct {
	data list.LinkedList[interface{}]
}

func NewQueue() *Queue {
	return &Queue{data: *list.NewLinkedList[interface{}]()}
}

func (queue *Queue) Size() int {
	return queue.data.Size()
}

func (queue *Queue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *Queue) Front() (interface{}, error) {
	return queue.data.Front()
}

func (queue *Queue) Back() (interface{}, error) {
	return queue.data.Back()
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.data.PushBack(value)
}

func (queue *Queue) Dequeue() (interface{}, error) {
	value, err := queue.data.Front()
	queue.data.PopFront()
	return value, err
}