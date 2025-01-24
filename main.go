package main

import (
	"fmt"
	"sstable/internal/trees"
)

func main() {
	index := trees.NewRBTree(10, 0)
	index.Insert(15, 0)
	index.Insert(20, 0)
	index.Insert(5, 0)
	fmt.Printf("%v %v %v %v", index.Root.LeftChild.Key, index.Root.Key, index.Root.RightChild.Key, index.Root.RightChild.RightChild.Key)
}
