package avl

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	value T
	parent *Node[T]
	left *Node[T]
	right *Node[T]
	height int
}

type AVL[T constraints.Ordered] struct {
	root *Node[T]
}

func (avl *AVL[T]) NewAVL(data []T) {
	for _, v := range data {
		avl.Insert(v)
	}
}

func (avl *AVL[T]) Search(value T) *Node[T]{
	node := avl.root
	for node != nil && node.value != value {
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

func (avl *AVL[T]) Insert(value T) *Node[T] {
	newNode := &Node[T]{value: value}
	var parent *Node[T]
	cur := avl.root
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
		avl.root = newNode
	} else if value < parent.value {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	avl.rebalance(newNode)
	return newNode
}

func (avl *AVL[T]) Delete(value T) bool {
	node := avl.Search(value)
	if node == nil {
		return false
	}
	actionNode := node.parent
	if node.left == nil {
		avl.transplant(node, node.right)
	} else if node.right == nil {
		avl.transplant(node, node.left)
	} else {
		successor := avl.Minimum(node.right)
		actionNode = successor.parent
		if successor.parent != node {
			avl.transplant(successor, successor.right)
			successor.right = node.right
			successor.right.parent = successor
		}
		avl.transplant(node, successor)
		successor.left = node.left
		successor.left.parent = successor
	}
	avl.rebalance(actionNode)
	return true
}

func (avl *AVL[T]) Minimum(node *Node[T]) *Node[T] {
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

func height[T constraints.Ordered](node *Node[T]) int {
	if node == nil {
		return -1
	}
	return node.height
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func updateHeight[T constraints.Ordered](node *Node[T]) {
	node.height = max(height(node.left), height(node.right)) + 1
}

func (avl *AVL[T]) leftRotate(x *Node[T]) {
	y := x.right
	y.parent = x.parent
	if y.parent == nil {
		avl.root = y
	} else {
		if y.parent.left == x {
			y.parent.left = y
		} else if y.parent.right == x {
			y.parent.right = y
		}
	}
	x.right = y.left
	if x.right != nil {
		x.right.parent = x
	}
	y.left = x
	x.parent = y
	updateHeight(x)
	updateHeight(y)
}

func (avl *AVL[T]) rightRotate(x *Node[T]) {
	y := x.left
	y.parent = x.parent
	if y.parent == nil {
		avl.root = y
	} else {
		if y.parent.left == x {
			y.parent.left = y
		} else if y.parent.right == x {
			y.parent.right = y
		}
	}
	x.left = y.right
	if x.left != nil {
		x.left.parent = x
	}
	y.right = x
	x.parent = y
	updateHeight(x)
	updateHeight(y)
}

func (avl *AVL[T]) rebalance(node *Node[T]) {
	for node != nil {
		updateHeight(node)
		if height(node.left) >= 2 + height(node.right) {
			if height(node.left.left) >= height(node.left.right) {
				avl.rightRotate(node)
			} else {
				avl.leftRotate(node.left)
				avl.rightRotate(node)
			}
		} else if height(node.right) >= 2 + height(node.left) {
			if height(node.right.right) >= height(node.right.left) {
				avl.leftRotate(node)
			} else {
				avl.rightRotate(node.right)
				avl.leftRotate(node)
			}
		}
		node = node.parent
	}
}

func (avl *AVL[T]) transplant(u, v *Node[T]) {
	if u.parent == nil {
		avl.root = v
	} else if u.parent.left == u {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}