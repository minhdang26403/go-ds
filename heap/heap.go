package heap

import (
	"fmt"

	"github.com/minhdang26403/algo-ds/utils"
	"golang.org/x/exp/constraints"
	"os"
)

type Heap[T constraints.Ordered] struct {
	data []T
	cmp utils.Comparator[T]
}

func NewHeap[T constraints.Ordered](data []T) *Heap[T] {
	return NewHeapWithComparator(data, utils.Less[T])
}

func NewHeapWithComparator[T constraints.Ordered](data []T, cmp utils.Comparator[T]) *Heap[T] {
	heap := &Heap[T]{data: make([]T, len(data)), cmp: cmp}
	n := copy(heap.data, data)
	if n != len(data) {
		fmt.Fprint(os.Stderr, "NewHeapWithComparator: copy error")
		return nil
	}
	heap.buildHeap()
	return heap
}

func (heap *Heap[T]) Top() (value T, err error) {
	if heap.IsEmpty() {
		return value, fmt.Errorf("Top: Attempt to access empty heap")
	}
	value = heap.data[0]
	return value, err
}

func (heap *Heap[T]) IsEmpty() bool { return len(heap.data) == 0 }

func (heap *Heap[T]) Size() int {
	return len(heap.data)
}

func (heap *Heap[T]) Push(value T) {
	heap.data = append(heap.data, value)
	heap.siftUp(heap.Size() - 1)
}

func (heap *Heap[T]) Pop() (T, error) {
	value, err := heap.Top()
	if err != nil {
		return value, err
	}
	n := heap.Size()
	heap.data[0], heap.data[n - 1] = heap.data[n - 1], heap.data[0]
	heap.data = heap.data[:n - 1]
	// Fix the heap invariant from the root
	heap.heapify(0)
	return value, err
}

func (heap *Heap[T]) String() string {
	return fmt.Sprint(heap.data)
}

/* Private helper methods */

func (heap *Heap[T]) getParent(i int) int { return (i - 1) / 2 }

func (heap *Heap[T]) getLeftChild(i int) int { return 2 * i + 1 }

func (heap *Heap[T]) getRightChild(i int) int { return 2 * i + 2 }

func (heap *Heap[T]) siftUp(i int) {
	parent := heap.getParent(i)
	for i > 0 && heap.cmp(heap.data[parent], heap.data[i]) {
		heap.data[parent], heap.data[i] = heap.data[i], heap.data[parent]
		i = parent
		parent = heap.getParent(i)
	}
}

func (heap *Heap[T]) buildHeap() {
	for i := (heap.Size() - 1) / 2; i >= 0; i-- {
		heap.heapify(i)
	}
}

func (heap *Heap[T]) heapify(i int) {
	largest := i
	left := heap.getLeftChild(i)
	right := heap.getRightChild(i)
	n := heap.Size()
	
	if left < n && heap.cmp(heap.data[largest], heap.data[left]) {
		largest = left
	}

	if right < n && heap.cmp(heap.data[largest], heap.data[right]) {
		largest = right
	}
	if largest != i {
		heap.data[largest], heap.data[i] = heap.data[i], heap.data[largest]
		heap.heapify(largest)
	}
}