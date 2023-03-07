package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	data := []int{83,86,77,15,93,35,11,92,49,21}
	bst := NewBST(data)
	
	root := bst.root
	if got := root.value; got != 83 {
		t.Errorf("Got %v expected %v", got, 83)
	}
	if got := root.left.value; got != 77 {
		t.Errorf("Got %v expected %v", got, 77)
	}
	if got := root.right.value; got != 86 {
		t.Errorf("Got %v expected %v", got, 86)
	}
	if got := bst.Search(93).left.value; got != 92 {
		t.Errorf("Got %v expected %v", got, 92)
	}
	if got := bst.Search(92).parent.value; got != 93 {
		t.Errorf("Got %v expected %v", got, 93)
	}
	if got := bst.Search(21).parent.value; got != 35 {
		t.Errorf("Got %v expected %v", got, 35)
	}
	if got := bst.Search(49).parent.value; got != 35 {
		t.Errorf("Got %v expected %v", got, 35)
	}
	if got := bst.Search(15).left.value; got != 11 {
		t.Errorf("Got %v expected %v", got, 11)
	}
	if got := bst.Search(15).right.value; got != 35 {
		t.Errorf("Got %v expected %v", got, 35)
	}
}

func TestDelete(t *testing.T) {
	data := []int{16,65,62,89,91,73,46,66,41,99}
	bst := NewBST(data)

	if got := bst.root.value; got != 16 {
		t.Errorf("Got %v expected %v", got, 16)
	}
	bst.Delete(16)
	if got := bst.root.value; got != 65 {
		t.Errorf("Got %v expected %v", got, 65)
	}
	if got := bst.Minimum(bst.root).value; got != 41 {
		t.Errorf("Got %v expected %v", got, 41)
	}
	if got := bst.Search(73).left.value; got != 66 {
		t.Errorf("Got %v expected %v", got, 66)
	}
	if got := bst.Search(73).parent.value; got != 89 {
		t.Errorf("Got %v expected %v", got, 89)
	}
	if got := bst.Search(91).left; got != nil {
		t.Errorf("Got %v expected %v", got, nil)
	}
}