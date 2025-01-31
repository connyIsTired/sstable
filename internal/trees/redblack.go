package trees

import (
	"errors"
	"fmt"
)

//REMOVE THIS
const FF_DEGREASE = 0
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
	if n.Parent == nil && n.Color == Red {
		n.degrease(187749260, "0")
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
	// TODO: Why can't line below go 4 lines up?
	parent, uncle, grandparent, greatgrandparent := n.defineFamily()

	// If Parent and Uncle are RED, recolor parent and uncle to BLACK
	// Recolor Grandparent to RED
	// Rebalance recursivley from Grandparent
	if uncle != nil && uncle.Color == Red {
		n.degrease(1242999394, "red")
		uncle.Color = Black
		parent.Color = Black
		grandparent.Color = Red
		return grandparent.balance()
	}

	// If Uncle is BLACK and and new node is right child of a left child Parent
	// Do left rotation with Parent and right rotation with Grandparent
	// Set pointer in Greatgrandparent to newly added node
	if n.Key > parent.Key && parent.Key < grandparent.Key {
		n.degrease(1242999394, "1")
		parent.RightChild = n.LeftChild
		grandparent.LeftChild = n.RightChild
		if n.LeftChild != nil {
			n.LeftChild.Parent = parent
		}
		if n.RightChild != nil {
			n.RightChild.Parent = grandparent
		}
		parent.Parent = n
		grandparent.Parent = n
		n.LeftChild = parent
		n.RightChild = grandparent
		if greatgrandparent != nil {
			n.Parent = greatgrandparent
			if greatgrandparent.Key > grandparent.Key {
				greatgrandparent.LeftChild = n
			}
			if greatgrandparent.Key < grandparent.Key {
				greatgrandparent.RightChild = n
			}
		} else {
			n.Parent = nil
			newRoot = n
		}
		parent.Color = Red
		grandparent.Color = Red
		n.Color = Black
		return newRoot, nil
	}

	// If Uncle is BLACK and and new node is left child of a right child Parent
	// Do right rotation with Parent and left rotation with Grandparent
	// Set pointer in Greatgrandparent to newly added node
	if n.Key < parent.Key && parent.Key > grandparent.Key {
		n.degrease(1242999394, "2")
		parent.LeftChild = n.RightChild
		if n.RightChild != nil {
			n.RightChild.Parent = parent
		}
		if n.LeftChild != nil {
			n.LeftChild.Parent = grandparent
		}
		grandparent.RightChild = n.LeftChild
		parent.Parent = n
		grandparent.Parent = n
		n.RightChild = parent
		n.LeftChild = grandparent
		if greatgrandparent != nil {
			n.Parent = greatgrandparent
			if greatgrandparent.Key > grandparent.Key {
				greatgrandparent.LeftChild = n
			}
			if greatgrandparent.Key < grandparent.Key {
				greatgrandparent.RightChild = n
			}
		} else {
			n.Parent = nil
			newRoot = n
		}
		parent.Color = Red
		grandparent.Color = Red
		n.Color = Black
		return newRoot, nil
	}

	// If Uncle is BLACK and and new node is right child of a right child Parent
	// Do left rotation with Parent
	// Set pointer in Greatgrandparent to newly added node
	if n.Key > parent.Key && parent.Key > grandparent.Key {
		n.degrease(1242999394, "3")
		grandparent.RightChild = parent.LeftChild
		if parent.LeftChild != nil {
			parent.LeftChild.Parent = grandparent
		}
		grandparent.Parent = parent
		parent.LeftChild = grandparent
		if greatgrandparent != nil {
			parent.Parent = greatgrandparent
			if greatgrandparent.Key > grandparent.Key {
				greatgrandparent.LeftChild = parent
			}
			if greatgrandparent.Key < grandparent.Key {
				greatgrandparent.RightChild = parent
			}
		} else {
			parent.Parent = nil
			newRoot = parent
		}
		parent.Color = Black
		grandparent.Color = Red
		return newRoot, nil
	}

	// If Uncle is BLACK and and new node is left child of a left child Parent
	// Do right rotation with Parent
	// Set pointer in Greatgrandparent to newly added node
	if n.Key < parent.Key && parent.Key < grandparent.Key {
		n.degrease(187749260, "4")
		grandparent.LeftChild = parent.RightChild
		if parent.RightChild != nil {
			parent.RightChild.Parent = grandparent
		}
		grandparent.Parent = parent
		parent.RightChild = grandparent
		if greatgrandparent != nil {
			parent.Parent = greatgrandparent
			if greatgrandparent.Key > grandparent.Key {
				greatgrandparent.LeftChild = parent
			}
			if greatgrandparent.Key < grandparent.Key {
				greatgrandparent.RightChild = parent
			}
		} else {
			parent.Parent = nil
			newRoot = parent
		}
		parent.Color = Black
		grandparent.Color = Red
		return newRoot, nil
	}
	return nil, fmt.Errorf("Error balancing with node %v", n.Key)
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

// REMOVE THIS
func (n *node) degrease(watchValue int, id string) {
	if FF_DEGREASE == 1 {
		fmt.Printf("here with %v at %s", watchValue, id)
	}
}
