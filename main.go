package main

import (
	"fmt"
	"sstable/internal/trees"
)

func main() {
	index := trees.NewRBTree(0, 0)
	index.Insert(1, 0)
	index.Insert(2, 0)
	index.Insert(3, 0)
	index.Insert(4, 0)

	fmt.Println(index.Root.Key)
	fmt.Println(index.Root.LeftChild.Key)
	fmt.Println(index.Root.RightChild.Key)
	fmt.Println(index.Root.RightChild.RightChild.Key)
	fmt.Println(index.Root.RightChild.LeftChild.Key)

	fmt.Println(index.Root.Color)
	fmt.Println(index.Root.LeftChild.Color)
	fmt.Println(index.Root.RightChild.Color)
	fmt.Println(index.Root.RightChild.RightChild.Color)
	fmt.Println(index.Root.RightChild.LeftChild.Color)
}
