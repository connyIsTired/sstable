package trees

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

type treeStateFn func(*node, *testing.T)

func TestCreateRBTreeRoot(t *testing.T) {
	randKey := rand.Int()
	randVal := rand.Int()
	rbTree := NewRBTree(randKey, randVal)
	if rbTree.Root.Key != randKey || rbTree.Root.Value != randVal {
		t.Fatalf("Root Node Key/Value was %v/%v, expected %v/%v", rbTree.Root.Key, rbTree.Root.Value, randKey, randVal)
	}
}

func TestCanInsertRootNodes(t *testing.T) {
	randKey := rand.Int()
	randVal := rand.Int()
	rbTree := NewRBTree(randKey, randVal)
	rbTree.Insert(randKey+10, randVal+10)
	rbTree.Insert(randKey-10, randVal-10)
	if rbTree.Root.RightChild.Key != randKey+10 || rbTree.Root.RightChild.Value != randVal+10 {
		t.Fatalf(
			"Root Right Child Key/Value was %v/%v, expected %v/%v",
			rbTree.Root.RightChild.Key,
			rbTree.Root.RightChild.Value,
			randKey+10,
			randVal+10)
	}
	if rbTree.Root.LeftChild.Key != randKey-10 || rbTree.Root.LeftChild.Value != randVal-10 {
		t.Fatalf(
			"Root Left Child Key/Value was %v/%v, expected %v/%v",
			rbTree.Root.LeftChild.Key,
			rbTree.Root.LeftChild.Value,
			randKey-10,
			randVal-10)
	}
}

func TestCanInsertGrandChildrenNodes(t *testing.T) {
	randKey := 100
	randVal := rand.Int()
	rbTree := NewRBTree(randKey, randVal)
	rbTree.Insert(randKey+10, randVal+10)
	rbTree.Insert(randKey+15, randVal+15)
	rbTree.Insert(randKey+5, randVal+5)
	if rbTree.Root.Key != randKey+10 {
		t.Fatalf("Root key was expected to be %v but was %v", randKey+10, rbTree.Root.Key)
	}
	if rbTree.Root.RightChild.Key != randKey+15 {
		t.Fatalf("Root key was expected to be %v but was %v", randKey+15, rbTree.Root.RightChild.Key)
	}
	if rbTree.Root.LeftChild.Key != randKey {
		t.Fatalf("Root key was expected to be %v but was %v", randKey, rbTree.Root.LeftChild.Key)
	}
	if rbTree.Root.LeftChild.RightChild.Key != randKey+5 {
		t.Fatalf("Root key was expected to be %v but was %v", randKey+5, rbTree.Root.LeftChild.RightChild.Key)
	}
}

func TestCanUpdateNodeValue(t *testing.T) {
	randKey := rand.Int()
	randVal := rand.Int()
	newVal := rand.Int()
	rbTree := NewRBTree(randKey, randVal)
	rbTree.Insert(randKey, newVal)

	if rbTree.Root.Value != newVal {
		t.Fatalf("Value in node was not changed. Was %v, expected %v", rbTree.Root.Value, newVal)
	}
}

func TestGetNodeReturnsValueIfPresent(t *testing.T) {
	var randVals []int
	rbTree := NewRBTree(0, 0)
	for i := 0; i < 10; i++ {
		randVal := rand.Int()
		randVals = append(randVals, randVal)
		rbTree.Insert(randVal, 0)
	}
	//randArrayIndex := rand.Intn(5)
	result := rbTree.GetNode(randVals[5])

	if result == nil {
		t.Fatalf("Expected to find node with key %v but returned nil\nValues in tree are %s\nValues in slice are %v", randVals[5], rbTree, randVals)
	}
}

func TestStringerIteratesCorrectly(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := "0 1 2 3 4 5 6 7 8"
	rbTree := NewRBTree(0, 0)
	for _, v := range input {
		rbTree.Insert(v, 0)
	}

	stringOutput := rbTree.String()

	result := strings.Compare(expected, stringOutput)
	if result != 0 {
		t.Fatalf("Slices are not equal.\nExpected %v\nGot %v", expected, stringOutput)
	}
}

func TestNoRedParentRedChildren(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	rbTree := NewRBTree(0, 0)
	for _, v := range input {
		rbTree.Insert(v, 0)
	}
	checkTreeState(rbTree.Root, noRedParentChild, t)
}

func TestNodesAreInCorrectOrder(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	rbTree := NewRBTree(0, 0)
	for _, v := range input {
		rbTree.Insert(v, 0)
	}
	checkTreeState(rbTree.Root, nodeAndChildrenInCorrectOrder, t)
}

func TestEachPathHasSameNumberOfNodes(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	rbTree := NewRBTree(0, 0)
	for _, v := range input {
		rbTree.Insert(v, 0)
	}
	ns := &[]*node{}
	getAllLeafNodes(rbTree.Root, ns)
	for i := range *ns {
		fmt.Println((*ns)[i].Key)
	}
}

func getAllLeafNodes(n *node, ns *[]*node) {
	if n.RightChild != nil {
		getAllLeafNodes(n.RightChild, ns)
	}
	if n.LeftChild != nil {
		getAllLeafNodes(n.LeftChild, ns)
	}
	if n.LeftChild == nil && n.RightChild == nil {
		*ns = append(*ns, n)
	}
}

func checkTreeState(node *node, fn treeStateFn, t *testing.T) {
	fn(node, t)
	if node.RightChild != nil {
		checkTreeState(node.RightChild, fn, t)
	}
	if node.LeftChild != nil {
		checkTreeState(node.LeftChild, fn, t)
	}
}

func noRedParentChild(n *node, t *testing.T) {
	if n.Color == Red {
		if n.RightChild != nil && n.RightChild.Color == Red {
			t.Fatal("Red child has red parent")
		}
		if n.LeftChild != nil && n.LeftChild.Color == Red {
			t.Fatal("Red child has red parent")
		}
	}
}

func nodeAndChildrenInCorrectOrder(n *node, t *testing.T) {
	if n.RightChild != nil && n.Key > n.RightChild.Key {
		t.Fatal("Right child less than parent")
	}
	if n.LeftChild != nil && n.Key < n.LeftChild.Key {
		t.Fatal("Left child greater than parent")
	}
}
