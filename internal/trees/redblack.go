package trees

import (
	"errors"
	"fmt"
)

const (
	Red color = iota
	Black
)

const (
	RightChildofRightChild childType = iota
	RightChildofLeftChild
	LeftChildofLeftChild
	LeftChildofRightChild
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
}

type nodeFamily struct {
	GreatGrandParent *node
	GrandParent      *node
	Parent           *node
	Unlce            *node
	IncomingNode     *node
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
	family := n.defineFamily()
	if n.Parent == nil && n.Color == Red {
		if n.Color == Red {
			n.Color = Black
			return nil, nil
		} else {
			return nil, fmt.Errorf("something strange happened")
		}
	}
	if n.Parent.Color == Black {
		return nil, nil
	}

	// If Parent and Uncle are RED, recolor parent and uncle to BLACK
	// Recolor Grandparent to RED
	// Rebalance recursivley from Grandparent
	if family.Unlce != nil && family.Unlce.Color == Red {
		family.Unlce.Color = Black
		family.Parent.Color = Black
		family.GrandParent.Color = Red
		return family.GrandParent.balance()
	}
	switch ct := family.getChildType(); ct {
	// If Uncle is BLACK and and new node is right child of a left child Parent
	// Do left rotation with Parent and right rotation with Grandparent
	// Set pointer in Greatgrandparent to newly added node
	case RightChildofLeftChild:
		return family.leftRightRotation(newRoot)

		// If Uncle is BLACK and and new node is left child of a right child Parent
		// Do right rotation with Parent and left rotation with Grandparent
		// Set pointer in Greatgrandparent to newly added node
	case LeftChildofRightChild:
		return family.rightLeftRotation(newRoot)

		// If Uncle is BLACK and and new node is right child of a right child Parent
		// Do left rotation with Parent
		// Set pointer in Greatgrandparent to newly added node
	case RightChildofRightChild:
		return family.leftRotation(newRoot)

		// If Uncle is BLACK and and new node is left child of a left child Parent
		// Do right rotation with Parent
		// Set pointer in Greatgrandparent to newly added node
	case LeftChildofLeftChild:
		return family.rightRotation(newRoot)
	default:
		return nil, fmt.Errorf("Error balancing with node %v", n.Key)
	}
}

func (n *node) defineFamily() *nodeFamily {
	family := nodeFamily{}
	if n.Parent != nil && n.Parent.Parent != nil {
		if n.Parent.Parent.Key > n.Parent.Key {
			family.Unlce = n.Parent.Parent.RightChild
		} else {
			family.Unlce = n.Parent.Parent.LeftChild
		}
	}
	if n.Parent != nil {
		family.Parent = n.Parent
		if n.Parent.Parent != nil {
			family.GrandParent = n.Parent.Parent
			if n.Parent.Parent.Parent != nil {
				family.GreatGrandParent = n.Parent.Parent.Parent
			}
		}
	}
	family.IncomingNode = n
	return &family
}

func (t *rbTree) GetNode(key int) *node {
	return t.Root.getNodeHelper(key)
}

func (n *node) getNodeHelper(key int) *node {
	if key == n.Key {
		return n
	}
	if key < n.Key && n.LeftChild != nil {
		return n.LeftChild.getNodeHelper(key)
	}
	if key > n.Key && n.RightChild != nil {
		return n.RightChild.getNodeHelper(key)
	}
	return nil
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

func (n *nodeFamily) getChildType() childType {
	if n.IncomingNode.Key < n.Parent.Key && n.Parent.Key < n.GrandParent.Key {
		return LeftChildofLeftChild
	}
	if n.IncomingNode.Key < n.Parent.Key && n.Parent.Key > n.GrandParent.Key {
		return LeftChildofRightChild
	}
	if n.IncomingNode.Key > n.Parent.Key && n.Parent.Key < n.GrandParent.Key {
		return RightChildofLeftChild
	}
	return RightChildofRightChild
}

func (n *nodeFamily) leftRightRotation(newRoot *node) (*node, error) {
	n.Parent.RightChild = n.IncomingNode.LeftChild
	n.GrandParent.LeftChild = n.IncomingNode.RightChild
	if n.IncomingNode.LeftChild != nil {
		n.IncomingNode.LeftChild.Parent = n.Parent
	}
	if n.IncomingNode.RightChild != nil {
		n.IncomingNode.RightChild.Parent = n.GrandParent
	}
	n.Parent.Parent = n.IncomingNode
	n.GrandParent.Parent = n.IncomingNode
	n.IncomingNode.LeftChild = n.Parent
	n.IncomingNode.RightChild = n.GrandParent
	if n.GreatGrandParent != nil {
		n.IncomingNode.Parent = n.GreatGrandParent
		if n.GreatGrandParent.Key > n.GrandParent.Key {
			n.GreatGrandParent.LeftChild = n.IncomingNode
		}
		if n.GreatGrandParent.Key < n.GrandParent.Key {
			n.GreatGrandParent.RightChild = n.IncomingNode
		}
	} else {
		n.IncomingNode.Parent = nil
		newRoot = n.IncomingNode
	}
	n.Parent.Color = Red
	n.GrandParent.Color = Red
	n.IncomingNode.Color = Black
	return newRoot, nil
}

func (n *nodeFamily) rightLeftRotation(newRoot *node) (*node, error) {
	n.Parent.LeftChild = n.IncomingNode.RightChild
	if n.IncomingNode.RightChild != nil {
		n.IncomingNode.RightChild.Parent = n.Parent
	}
	if n.IncomingNode.LeftChild != nil {
		n.IncomingNode.LeftChild.Parent = n.GrandParent
	}
	n.GrandParent.RightChild = n.IncomingNode.LeftChild
	n.Parent.Parent = n.IncomingNode
	n.GrandParent.Parent = n.IncomingNode
	n.IncomingNode.RightChild = n.Parent
	n.IncomingNode.LeftChild = n.GrandParent
	if n.GreatGrandParent != nil {
		n.IncomingNode.Parent = n.GreatGrandParent
		if n.GreatGrandParent.Key > n.GrandParent.Key {
			n.GreatGrandParent.LeftChild = n.IncomingNode
		}
		if n.GreatGrandParent.Key < n.GrandParent.Key {
			n.GreatGrandParent.RightChild = n.IncomingNode
		}
	} else {
		n.IncomingNode.Parent = nil
		newRoot = n.IncomingNode
	}
	n.Parent.Color = Red
	n.GrandParent.Color = Red
	n.IncomingNode.Color = Black
	return newRoot, nil
}

func (n *nodeFamily) leftRotation(newRoot *node) (*node, error) {
	n.GrandParent.RightChild = n.Parent.LeftChild
	if n.Parent.LeftChild != nil {
		n.Parent.LeftChild.Parent = n.GrandParent
	}
	n.GrandParent.Parent = n.Parent
	n.Parent.LeftChild = n.GrandParent
	if n.GreatGrandParent != nil {
		n.Parent.Parent = n.GreatGrandParent
		if n.GreatGrandParent.Key > n.GrandParent.Key {
			n.GreatGrandParent.LeftChild = n.Parent
		}
		if n.GreatGrandParent.Key < n.GrandParent.Key {
			n.GreatGrandParent.RightChild = n.Parent
		}
	} else {
		n.Parent.Parent = nil
		newRoot = n.Parent
	}
	n.Parent.Color = Black
	n.GrandParent.Color = Red
	return newRoot, nil
}

func (n *nodeFamily) rightRotation(newRoot *node) (*node, error) {
	n.GrandParent.LeftChild = n.Parent.RightChild
	if n.Parent.RightChild != nil {
		n.Parent.RightChild.Parent = n.GrandParent
	}
	n.GrandParent.Parent = n.Parent
	n.Parent.RightChild = n.GrandParent
	if n.GreatGrandParent != nil {
		n.Parent.Parent = n.GreatGrandParent
		if n.GreatGrandParent.Key > n.GrandParent.Key {
			n.GreatGrandParent.LeftChild = n.Parent
		}
		if n.GreatGrandParent.Key < n.GrandParent.Key {
			n.GreatGrandParent.RightChild = n.Parent
		}
	} else {
		n.Parent.Parent = nil
		newRoot = n.Parent
	}
	n.Parent.Color = Black
	n.GrandParent.Color = Red
	return newRoot, nil
}
