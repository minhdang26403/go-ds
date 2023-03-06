package queue

import "testing"

func TestGeneral(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	if got := q.Size(); got != 3 {
		t.Errorf("Got %v expected %v", got, 3)
	}

	if got, err := q.Front(); got != 1 || err != nil {
		t.Errorf("Got %v expected %v", got, 1)
	}

	if got, err := q.Back(); got != 3 || err != nil {
		t.Errorf("Got %v expected %v", got, 3)
	}

	q.Dequeue()
	if got, err := q.Front(); got != 2 || err != nil {
		t.Errorf("Got %v expected %v", got, 2)
	}
}