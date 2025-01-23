package redblack

const (
	Red Color = iota
	Black
)

type Color int

type Node struct {
	Key        int
	Value      int
	Parent     *Node
	LeftChild  *Node
	RightChild *Node
	Color      Color
}

func NewNode() *Node {
	return &Node{Color: Red}
}
