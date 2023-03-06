package heap

import (
	"testing"
	"github.com/minhdang26403/algo-ds/utils"
)

func TestMaxHeap(t *testing.T) {
	data := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap := NewHeap(data)
	
	if got := heap.String(); got != "[16 14 10 8 7 9 3 2 4 1]" {
		t.Errorf("Got %v expected %v", got, "[16 14 10 8 7 9 3 2 4 1]")
	}

	if got, _ := heap.Top(); got != 16 {
		t.Errorf("Got %v expected %v", got, 16)
	}
	heap.Push(17)
	if got, _ := heap.Top(); got != 17 {
		t.Errorf("Got %v expected %v", got, 17)
	}
}

func TestPushMaxHeap(t *testing.T) {
	heap := NewHeap([]int{})
	for i := 0; i < 11; i++ {
		heap.Push(i)
	}

	if got, _ := heap.Top(); got != 10 {
		t.Errorf("Got %v expected %v", got, 10)
	}

	if got, _ := heap.Pop(); got != 10 {
		t.Errorf("Got %v expected %v", got, 10)
	}

	if got, _ := heap.Top(); got != 9 {
		t.Errorf("Got %v expected %v", got, 9)
	}

	heap.Push(20)
	if got, _ := heap.Top(); got != 20 {
		t.Errorf("Got %v expected %v", got, 20)
	}
	for i := 0; i < 5; i++ {
		heap.Pop()
	}
	if got, _ := heap.Top(); got != 5 {
		t.Errorf("Got %v expected %v", got, 5)
	}
}

func TestMinHeap(t *testing.T) {
	data := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap := NewHeapWithComparator(data, utils.Greater[int])
	if got, _ := heap.Top(); got != 1 {
		t.Errorf("Got %v expected %v", got, 1)
	}
	heap.Push(-1)
	if got, _ := heap.Top(); got != -1 {
		t.Errorf("Got %v expected %v", got, -1)
	}
	if got, _ := heap.Pop(); got != -1 {
		t.Errorf("Got %v expected %v", got, -1)
	}
	if got, _ := heap.Pop(); got != 1 {
		t.Errorf("Got %v expected %v", got, 1)
	}
	if got, _ := heap.Pop(); got != 2 {
		t.Errorf("Got %v expected %v", got, 2)
	}
	if got, _ := heap.Top(); got != 3 {
		t.Errorf("Got %v expected %v", got, 3)
	}
}

func TestEmptyHeap(t *testing.T) {
	heap := NewHeap([]int{})
	if _, err := heap.Top(); err == nil {
		t.Errorf("Get no error when accessing empty heap")
	}
}