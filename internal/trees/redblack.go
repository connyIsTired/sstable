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
	return t.Root.insert(key, value)
}

func (parentNode *node) insert(key int, value int) error {
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
		return nil
	}
	return errors.New("Node could not be inserted")
}

func (n *node) balance() error {
	if n.Parent.Color == Black {
		return nil
	}

	return errors.New("Could not balance")
}

func (n *node) findUncleColor() color {
	if n.ChildType == LeftChild {
		return n.Parent.RightChild.Color
	}
	return n.Parent.LeftChild.Color
}
