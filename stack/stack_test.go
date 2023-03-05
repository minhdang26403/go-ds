package stack

import (
	"testing"
)

func TestBasic(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(2)
	stack.Push(6)
	stack.Push(4)
	if got := stack.Size(); got != 3 {
		t.Errorf("Got %v expected %v", got, 3)
	}
}

func TestStackPush(t *testing.T) {
	stack := NewStack[int]()
	if got := stack.IsEmpty(); got != true {
		t.Errorf("Got %v expected %v", got, true)
	}
	stack.Push(2)
	stack.Push(6)
	stack.Push(4)

	if got := stack.IsEmpty(); got != false {
		t.Errorf("Got %v expected %v", got, false)
	}
	if got := stack.Size(); got != 3 {
		t.Errorf("Got %v expected %v", got, 3)
	}
	if got, err := stack.Top(); got != 4 || err != nil {
		t.Errorf("Got %v expected %v", got, 3)
	}
}

func TestStackTop(t *testing.T) {
	stack := NewStack[int]()
	if _, err := stack.Top(); err != nil {
		t.Errorf("Got value from an empty stack")
	}
	stack.Push(2)
	stack.Push(5)
	stack.Push(10)
	if got, err := stack.Top(); got != 10 || err != nil {
		t.Errorf("Got %v expected %v", got, 3)
	}
}

func TestStackPop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(4)
	stack.Push(2)
	stack.Push(6)
	stack.Pop()
	if got, err := stack.Top(); got != 2 || err != nil {
		t.Errorf("Got %v expected %v", got, 6)
	}
	if got, err := stack.Pop(); got != 2 || err != nil {
		t.Errorf("Got %v expected %v", got, 2)
	}
	if got, err := stack.Pop(); got != 4 || err != nil {
		t.Errorf("Got %v expected %v", got, 1)
	}
	if _, err := stack.Pop(); err != nil {
		t.Errorf("Pop from an empty stack")
	}
	if got := stack.IsEmpty(); got != true {
		t.Errorf("Got %v expected %v", got, true)
	}
}

func benchmarkPush(b *testing.B, stack *Stack[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Push(n)
		}
	}
}

func benchmarkPop[T any](b *testing.B, stack *Stack[T], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Pop()
		}
	}
}

func BenchmarkArrayStackPop100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := NewStack[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkArrayStackPop1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := NewStack[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkArrayStackPop10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := NewStack[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkArrayStackPop100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := NewStack[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkArrayStackPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := NewStack[int]()
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkArrayStackPush1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := NewStack[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkArrayStackPush10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := NewStack[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkArrayStackPush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := NewStack[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}
