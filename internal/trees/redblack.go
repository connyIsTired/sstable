package trees

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
	node := &node{Color: Red, Key: key, Value: value}
	return &rbTree{Root: node}
}

func (t *rbTree) Insert(key int, value int) error {
	return t.Root.insert(key, value)
}

func (n *node) insert(key int, value int) error {
	if key > n.Key {
		if n.RightChild == nil {
			newnode := &node{Key: key, Value: value}
			n.RightChild = newnode
			newnode.Parent = n
			return nil
		}
		n.RightChild.insert(key, value)
		return nil
	}
	if key < n.Key {
		if n.LeftChild == nil {
			newnode := &node{Key: key, Value: value}
			n.LeftChild = newnode
			newnode.Parent = n
			return nil
		}
		n.LeftChild.insert(key, value)
		return nil
	}
	return nil
}
