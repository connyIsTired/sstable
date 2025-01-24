package main

import (
	"fmt"
	"sstable/internal/trees"
)

func main() {
	index := trees.NewRBTree(110, 0)
	err := index.Insert(120, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = index.Insert(115, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v,%v,%v", index.Root.Key, index.Root.LeftChild.Key, index.Root.RightChild.Key)
}
