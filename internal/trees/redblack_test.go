package trees

import (
	"math/rand"
	"testing"
)

func TestCreateRBTreeRoot(t *testing.T) {
	randKey := rand.Int()
	randVal := rand.Int()
	index := NewRBTree(randKey, randVal)
	if index.Root.Key != randKey || index.Root.Value != randVal {
		t.Fatalf("Root Node Key/Value was %v/%v, expected %v/%v", index.Root.Key, index.Root.Value, randKey, randVal)
	}
}

func TestCanInsertRootNodes(t *testing.T) {
	randKey := rand.Int()
	randVal := rand.Int()
	index := NewRBTree(randKey, randVal)
	index.Insert(randKey+10, randVal+10)
	index.Insert(randKey-10, randVal-10)
	if index.Root.RightChild.Key != randKey+10 || index.Root.RightChild.Value != randVal+10 {
		t.Fatalf(
			"Root Right Child Key/Value was %v/%v, expected %v/%v",
			index.Root.RightChild.Key,
			index.Root.RightChild.Value,
			randKey+10,
			randVal+10)
	}
	if index.Root.LeftChild.Key != randKey-10 || index.Root.LeftChild.Value != randVal-10 {
		t.Fatalf(
			"Root Left Child Key/Value was %v/%v, expected %v/%v",
			index.Root.LeftChild.Key,
			index.Root.LeftChild.Value,
			randKey-10,
			randVal-10)
	}
}

func TestCanInsertGrandChildrenNodes(t *testing.T) {
	randKey := rand.Int()
	randVal := rand.Int()
	index := NewRBTree(randKey, randVal)
	index.Insert(randKey+10, randVal+10)
	index.Insert(randKey+15, randVal+15)
	index.Insert(randKey+5, randVal+5)
	childNode := index.Root.RightChild
	if childNode.RightChild.Key != randKey+15 || childNode.RightChild.Value != randVal+15 {
		t.Fatalf(
			"Child Right Child Key/Value was %v/%v, expected %v/%v",
			childNode.RightChild.Key,
			childNode.RightChild.Value,
			randKey+15,
			randVal+15)
	}
	if childNode.LeftChild.Key != randKey+5 || childNode.LeftChild.Value != randVal+5 {
		t.Fatalf(
			"Child Left Child Key/Value was %v/%v, expected %v/%v",
			childNode.LeftChild.Key,
			childNode.LeftChild.Value,
			randKey+5,
			randVal+5)
	}
}

func TestCanUpdateNodeValue(t *testing.T) {
	randKey := rand.Int()
	randVal := rand.Int()
	newVal := rand.Int()
	index := NewRBTree(randKey, randVal)
	index.Insert(randKey, newVal)

	if index.Root.Value != newVal {
		t.Fatalf("Value in node was not changed. Was %v, expected %v", index.Root.Value, newVal)
	}
}
