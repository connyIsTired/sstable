package trees

import (
	"errors"
	"fmt"
)

const (
	Red color = iota
	Black
)

type color int

type rbTree struct {
	Root *node
}

type node struct {
	Key        int
	Value      int
	Parent     *node
	LeftChild  *node
	RightChild *node
	Color      color
}

func NewRBTree(key int, value int) *rbTree {
	node := &node{Color: Black, Key: key, Value: value}
	return &rbTree{Root: node}
}

func (t *rbTree) Insert(key int, value int) error {
	newRoot, err := t.Root.insert(key, value)
	if err != nil {
		return err
	}
	if newRoot != nil {
		t.Root = newRoot
	}
	return nil
}

func (parentNode *node) insert(key int, value int) (*node, error) {
	if key > parentNode.Key {
		if parentNode.RightChild == nil {
			newnode := &node{Key: key, Value: value, Color: Red}
			parentNode.RightChild = newnode
			newnode.Parent = parentNode
			return newnode.balance()
		}
		return parentNode.RightChild.insert(key, value)
	}
	if key < parentNode.Key {
		if parentNode.LeftChild == nil {
			newnode := &node{Key: key, Value: value, Color: Red}
			parentNode.LeftChild = newnode
			newnode.Parent = parentNode
			return newnode.balance()
		}
		return parentNode.LeftChild.insert(key, value)
	}
	if key == parentNode.Key {
		parentNode.Value = value
		return nil, nil
	}
	return nil, errors.New("Node could not be inserted")
}

func (n *node) balance() (*node, error) {
	var newRoot *node
	if n.Parent.Color == Black {
		return nil, nil
	}
	parent, uncle, grandparent, greatgrandparent := n.defineFamily()

	if uncle == nil || uncle.Color == Black {
		if n.Key < parent.Key {
			if grandparent.Parent == nil {
				newRoot = parent
			}
			parent.Parent = grandparent.Parent
			greatgrandparent.LeftChild = parent
			grandparent.LeftChild = parent.RightChild
			parent.RightChild = grandparent
			grandparent.Parent = parent
			parent.Color = Black
			grandparent.Color = Red
			return newRoot, nil
		}
		if grandparent.Parent == nil {
			newRoot = parent
		}
		parent.LeftChild = grandparent
		grandparent.Parent = n.Parent
		return newRoot, nil
	}
	parent.Color = Black
	uncle.Color = Black
	if grandparent.Parent == nil {
		grandparent.Color = Black
		return nil, nil
	}
	grandparent.Color = Red
	return grandparent.balance()
}

func (n *node) defineFamily() (parent, uncle, grandparent, greatgrandparent *node) {
	if n.Parent.Parent.Key > n.Parent.Key {
		uncle = n.Parent.Parent.RightChild
	} else {
		uncle = n.Parent.Parent.LeftChild
	}
	parent = n.Parent
	grandparent = parent.Parent
	greatgrandparent = grandparent.Parent
	return
}
func (t *rbTree) String() string {
	if t.Root == nil {
		return "[]"
	}
	return t.Root.stringHelper()
}

func (n *node) stringHelper() string {
	if n == nil {
		return ""
	}
	var result string
	result = fmt.Sprintf("%v", n.Key)
	if n.LeftChild != nil {
		result = fmt.Sprintf("%v %v", n.LeftChild.stringHelper(), result)
	}
	if n.RightChild != nil {
		result = fmt.Sprintf("%v %v", result, n.RightChild.stringHelper())
	}
	return result
}
