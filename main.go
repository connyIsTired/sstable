package main

import (
	"fmt"
	"sstable/internal/trees"
)

func main() {
	index := trees.NewRBTree(10, 0)
	err := index.Insert(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = index.Insert(1, 0)
	if err != nil {
		fmt.Println(err)
	}
}
