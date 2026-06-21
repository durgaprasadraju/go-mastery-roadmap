// Package bst implements a binary search tree from scratch.
package bst

import "errors"

var ErrNotFound = errors.New("key not found")

// Node is a BST node.
type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}

// BST is a binary search tree.
type BST struct {
	root *Node
	size int
}

// New creates an empty BST.
func New() *BST {
	return &BST{}
}

// Insert adds a key-value pair. O(h) where h is height.
func (t *BST) Insert(key int, value string) {
	t.root = t.insert(t.root, key, value)
}

func (t *BST) insert(n *Node, key int, value string) *Node {
	if n == nil {
		t.size++
		return &Node{Key: key, Value: value}
	}
	if key < n.Key {
		n.Left = t.insert(n.Left, key, value)
	} else if key > n.Key {
		n.Right = t.insert(n.Right, key, value)
	} else {
		n.Value = value // update
	}
	return n
}

// Search finds a value by key. O(h).
func (t *BST) Search(key int) (string, error) {
	n := t.search(t.root, key)
	if n == nil {
		return "", ErrNotFound
	}
	return n.Value, nil
}

func (t *BST) search(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if key == n.Key {
		return n
	}
	if key < n.Key {
		return t.search(n.Left, key)
	}
	return t.search(n.Right, key)
}

// InOrderTraversal returns keys in sorted order. O(n).
func (t *BST) InOrderTraversal() []int {
	out := make([]int, 0, t.size)
	var walk func(*Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		walk(n.Left)
		out = append(out, n.Key)
		walk(n.Right)
	}
	walk(t.root)
	return out
}

// Size returns node count.
func (t *BST) Size() int { return t.size }

// Min returns the minimum key. O(h).
func (t *BST) Min() (int, error) {
	if t.root == nil {
		return 0, ErrNotFound
	}
	n := t.root
	for n.Left != nil {
		n = n.Left
	}
	return n.Key, nil
}

// Delete removes a key. O(h).
func (t *BST) Delete(key int) error {
	var err error
	t.root, err = t.delete(t.root, key)
	return err
}

func (t *BST) delete(n *Node, key int) (*Node, error) {
	if n == nil {
		return nil, ErrNotFound
	}
	if key < n.Key {
		var err error
		n.Left, err = t.delete(n.Left, key)
		return n, err
	}
	if key > n.Key {
		var err error
		n.Right, err = t.delete(n.Right, key)
		return n, err
	}
	t.size--
	if n.Left == nil {
		return n.Right, nil
	}
	if n.Right == nil {
		return n.Left, nil
	}
	succ := n.Right
	for succ.Left != nil {
		succ = succ.Left
	}
	n.Key, n.Value = succ.Key, succ.Value
	var err error
	n.Right, err = t.delete(n.Right, succ.Key)
	return n, err
}
