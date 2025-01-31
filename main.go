package main

import (
	"fmt"
	"sstable/internal/trees"
)

func main() {
	x := []int{1726, 1479, 1331, 7289, 2773, 2101, 3175, 7436, 4742, 5449}
	index := trees.NewRBTree(0, 0)
	for i := range x {
		index.Insert(x[i], 0)
		//fmt.Printf("inserting %v\n", x[i])
		//fmt.Println(index)
		//fmt.Printf("root is %v\n", index.Root.Key)
	}
	fmt.Println(index.GetNode(2101))

}
