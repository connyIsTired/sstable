package trees

import (
	"errors"
)

const (
	Red color = iota
	Black
)

const (
	LeftChild childType = iota
	RightChild
)

type color int
type childType int

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
	ChildType  childType
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
			newnode.ChildType = RightChild
			return newnode.balance()
		}
		return parentNode.RightChild.insert(key, value)
	}
	if key < parentNode.Key {
		if parentNode.LeftChild == nil {
			newnode := &node{Key: key, Value: value, Color: Red}
			parentNode.LeftChild = newnode
			newnode.Parent = parentNode
			newnode.ChildType = LeftChild
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
	parent, uncle, grandparent := n.defineFamily()

	if uncle == nil || uncle.Color == Black {
		if n.ChildType == LeftChild {
			if grandparent.Parent == nil {
				newRoot = parent
			}
			parent.Parent = grandparent.Parent
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
	return nil, errors.New("Could not balance")
}

func (n *node) findUncle() *node {
	if n.Parent.ChildType == LeftChild {
		return n.Parent.Parent.RightChild
	}
	return n.Parent.Parent.LeftChild
}

func (n *node) defineFamily() (parent, uncle, grandparent *node) {
	if n.Parent.ChildType == LeftChild {
		uncle = n.Parent.Parent.RightChild
	} else {
		uncle = n.Parent.Parent.LeftChild
	}
	parent = n.Parent
	grandparent = n.Parent.Parent
	return
}
