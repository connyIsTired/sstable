package main

import (
	"fmt"
	"sstable/internal/trees"
)

func main() {
	index := trees.NewRBTree(10, 0)
	index.Insert(5, 0)
	index.Insert(15, 0)
	index.Insert(13, 0)
	index.Insert(17, 0)
	index.Insert(3, 0)
	index.Insert(7, 0)
	index.Insert(16, 0)
	fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v",
		index.Root.Color,
		index.Root.LeftChild.Color,
		index.Root.LeftChild.LeftChild.Color,
		index.Root.LeftChild.RightChild.Color,
		index.Root.RightChild.Color,
		index.Root.RightChild.LeftChild.Color,
		index.Root.RightChild.RightChild.Color,
		index.Root.RightChild.RightChild.LeftChild.Color,
	)
}
