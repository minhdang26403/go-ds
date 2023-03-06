package linkedlist

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	value T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// `NewLinkedLit` instantiates a new list and adds the passed values, if any, to the list
func NewLinkedList[T any](values ...T) *LinkedList[T] {
	list := &LinkedList[T]{}
	for _, value := range values {
		newNode := &Node[T]{value: value, prev: list.tail}
		if list.size == 0 {
			list.head = newNode
		} else {
			list.tail.next = newNode
		}
		list.tail = newNode
		list.size++
	}
	return list
}

/** Element access */
// `Front` returns a pointer to the value of the first element in the linked list
func (list *LinkedList[T]) Front() T {
	return list.head.value
}

// `Front` returns a pointer to the value of the last element in the linked list
func (list *LinkedList[T]) Back() T {
	return list.tail.value
}

// `Begin` returns a pointer to the first element of the list
func (list *LinkedList[T]) Begin() *Node[T] {
	return list.head
}

// `Begin` returns a pointer to the last element of the list
func (list *LinkedList[T]) End() *Node[T] {
	return list.tail
}

// `Size` returns the size of the linked list
func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.Size() == 0
}

// `String` returns a string representation of the linked list
func (list *LinkedList[T]) String() string {
	if list.head == nil {
		return "[]"
	}
	var s strings.Builder
	s.WriteByte('[')
	cur := list.head
	for cur.next != nil {
		s.WriteString(fmt.Sprintf("%v", cur.value))
		s.WriteString(", ")
		cur = cur.next
	}
	s.WriteString(fmt.Sprintf("%v", cur.value))
	s.WriteByte(']')
	return s.String()
}

// `Insert` inserts a `value` into `index` position of the list
func (list *LinkedList[T]) Insert(index int, value T) error {
	if index < 0 || index > list.Size() {
		return fmt.Errorf("Insert: Invalid index")
	}
	newNode := &Node[T]{value: value}
	if index == 0 {
		newNode.next = list.head
		if list.head != nil {
			list.head.prev = newNode
		}
		list.head = newNode
	} else {
		prevNode, _ := list.find(index - 1)
		newNode.next = prevNode.next
		newNode.prev = prevNode
		prevNode.next = newNode
	}
	if index == list.size {
		list.tail = newNode
	} else {
		newNode.next.prev = newNode
	}
	list.size++
	return nil
}

// `Erase` removes an element at position `index` in the list
func (list *LinkedList[T]) Erase(index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Invalid index")
	}

	if index == 0 {
		list.head = list.head.next
		if list.head == nil {
			list.tail = nil
		}
	} else {
		prevNode, _ := list.find(index - 1)
		prevNode.next = prevNode.next.next
		if index == list.size - 1 {
			list.tail = prevNode
		} else {
			prevNode.next.prev = prevNode
		}
	}
	list.size--
	return nil
}

func (list *LinkedList[T]) PushBack(value T) {
	list.Insert(list.size, value)
}

func (list *LinkedList[T]) PushFront(value T) {
	list.Insert(0, value)
}

func (list *LinkedList[T]) PopBack() error {
	return list.Erase(list.size - 1)
}

func (list *LinkedList[T]) PopFront() error {
	return list.Erase(0)
}


func (list *LinkedList[T]) Reverse() {
	cur := list.head
	for cur != nil {
		next := cur.next
		cur.next = cur.prev
		cur.prev = next
		cur = next
	}
	list.head, list.tail = list.tail, list.head
}

func (list *LinkedList[T]) find(index int) (node *Node[T], err error) {
	if index < 0 || index >= list.size {
		return node, fmt.Errorf("find: Invalid index")
	}

	node = list.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, err
}