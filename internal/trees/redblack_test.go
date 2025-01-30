package trees

import (
	"math/rand"
	"strings"
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

	if index.Root.Key != randKey+10 {
		t.Fatalf("Root key was expected to be %v but was %v", randKey+10, index.Root.Key)
	}
	if index.Root.RightChild.Key != randKey+15 {
		t.Fatalf("Root key was expected to be %v but was %v", randKey+15, index.Root.RightChild.Key)
	}
	if index.Root.LeftChild.Key != randKey {
		t.Fatalf("Root key was expected to be %v but was %v", randKey, index.Root.LeftChild.Key)
	}
	if index.Root.LeftChild.RightChild.Key != randKey+5 {
		t.Fatalf("Root key was expected to be %v but was %v", randKey+5, index.Root.LeftChild.RightChild.Key)
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

func TestStringerIteratesCorrectly(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := "0 1 2 3 4 5 6 7 8"
	index := NewRBTree(0, 0)

	for _, v := range input {
		index.Insert(v, 0)
	}

	stringOutput := index.String()

	result := strings.Compare(expected, stringOutput)
	if result != 0 {
		t.Fatalf("Slices are not equal.\nExpected %v\nGot %v", expected, stringOutput)
	}
}
