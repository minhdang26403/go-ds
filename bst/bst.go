package bst

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	value  T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

func (node *Node[T]) Value() T {
	return node.value
}

func (node *Node[T]) Parent() *Node[T] {
	return node.parent
}

func (node *Node[T]) Left() *Node[T] {
	return node.left
}

func (node *Node[T]) Right() *Node[T] {
	return node.right
}

type BST[T constraints.Ordered] struct {
	root *Node[T]
}

func NewBST[T constraints.Ordered](values []T) *BST[T] {
	bst := &BST[T]{}
	for _, value := range values {
		bst.Insert(value)
	}
	return bst
}

func (bst *BST[T]) GetRoot() *Node[T] {
	return bst.root
}

func (bst *BST[T]) Search(value T) *Node[T] {
	node := bst.root
	for node != nil && node.value != value {
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

func (bst *BST[T]) Minimum(node *Node[T]) *Node[T] {
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BST[T]) Maximum(node *Node[T]) *Node[T] {
	for node != nil && node.right != nil {
		node = node.right
	}
	return node
}

func (bst *BST[T]) Successor(node *Node[T]) *Node[T] {
	if node.right != nil {
		return bst.Minimum(node.right)
	}
	for node.parent != nil && node == node.parent.right {
		node = node.parent
	}
	return node.parent
}

func (bst *BST[T]) Predecessor(node *Node[T]) *Node[T] {
	if node.left != nil {
		return bst.Maximum(node.left)
	}
	for node.parent != nil && node == node.parent.left {
		node = node.parent
	}
	return node.parent
}

func (bst *BST[T]) Insert(value T) *Node[T] {
	newNode := &Node[T]{value: value}
	var parent *Node[T]
	cur := bst.root
	for cur != nil {
		parent = cur
		if value < cur.value {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	newNode.parent = parent
	if parent == nil {
		bst.root = newNode
	} else if value < parent.value {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	return newNode
}

func (bst *BST[T]) Delete(value T) bool {
	node := bst.Search(value)
	if node == nil {
		return false
	}
	if node.left == nil {
		bst.transplant(node, node.right)
	} else if node.right == nil {
		bst.transplant(node, node.left)
	} else {
		successor := bst.Minimum(node.right)
		if successor.parent != node {
			bst.transplant(successor, successor.right)
			successor.right = node.right
			successor.right.parent = successor
		}
		bst.transplant(node, successor)
		successor.left = node.left
		successor.left.parent = successor
	}
	return true
}

func (bst *BST[T]) GetValues() []T {
	var data []T
	bst.inorderTraversal(bst.root, data)
	return data
}

func (bst *BST[T]) inorderTraversal(root *Node[T], data []T) {
	if root == nil {
		return
	}
	bst.inorderTraversal(root.left, data)
	data = append(data, root.value)
	bst.inorderTraversal(root.right, data)
}

func (bst *BST[T]) transplant(u, v *Node[T]) {
	if u.parent == nil {
		bst.root = v
	} else if u.parent.left == u {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}
