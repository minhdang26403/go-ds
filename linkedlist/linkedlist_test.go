package linkedlist

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	list := NewLinkedList(1, 2, 3, 4)
	if got := list.String(); got != "[1, 2, 3, 4]" {
		t.Errorf("Got %v expected %v", got, "[1, 2, 3, 4]")
	}
	emptyList := NewLinkedList[int]()
	if got := emptyList.String(); got != "[]" {
		t.Errorf("Got %v expected %v", got, "[]")
	}
}

func TestInsert(t *testing.T) {
	list := NewLinkedList[string]()
	list.Insert(0, "yyang")
	list.Insert(0, "yoonalim")
	if got := list.String(); got != "[yoonalim, yyang]" {
		t.Errorf("Got %v expected %v", got, "[yoonalim, yyang]")
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkedList[string]()
	list.Insert(0, "abc")
	list.Insert(1, "def")
	list.Insert(1, "ghi")
	list.Insert(0, "xyz")
	
	list.Erase(0)
	if got := list.Front(); got != "abc" {
		t.Errorf("Got %v expected %v", got, "abc")
	}
	
	list.Erase(0)
	if got := list.Front(); got != "ghi" {
		t.Errorf("Got %v expected %v", got, "ghi")
	}

	list.Erase(1)
	if got := list.String(); got != "[ghi]" {
		t.Errorf("Got %v expected %v", got, "[ghi]")
	}
}

func TestReverse(t *testing.T) {
	list := NewLinkedList[string]()
	list.Reverse()
	if got := list.String(); got != "[]" {
		t.Errorf("Got %v expected %v", got, "[]")
	}

	list.Insert(0, "a")
	list.Insert(0, "b")
	if got := list.String(); got != "[b, a]" {
		t.Errorf("Got %v expected %v", got, "[b, a]")
	}

	list.Insert(2, "c")
	list.Reverse()
	if got := list.String(); got != "[c, a, b]" {
		t.Errorf("Got %v expected %v", got, "[c, a, b]")
	}

	list.Erase(1)
	list.Reverse()
	if got := list.String(); got != "[b, c]" {
		t.Errorf("Got %v expected %v", got, "[b, c]")
	}
}