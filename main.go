package main

import (
	"fmt"
	"sstable/internal/trees"
)

func main() {
	index := trees.NewRBTree(100, 0)
	index.Insert(110, 0)
	index.Insert(115, 0)
	index.Insert(105, 0)

	fmt.Println(index.Root.Key)
	fmt.Println(index.Root.LeftChild.Key)
	fmt.Println(index.Root.LeftChild.RightChild.Key)
	fmt.Println(index.Root.RightChild.Key)
}
