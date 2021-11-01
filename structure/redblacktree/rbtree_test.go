package redblacktree

import (
	"fmt"
	"testing"
)

func compare(tree *RBTree, preVal []int, preRed []bool, inVal []int, inRed []bool) bool {
	var ret = false
	var pre = preOrder(tree.Root)
	var in = inOrder(tree.Root)
	if len(pre) != len(preVal) || len(in) != len(inVal) {
		return ret
	}
	for i := 0; i < len(pre); i++ {
		if pre[i].Val != preVal[i] || pre[i].isRed != preRed[i] {
			fmt.Printf("Error at pre %d", i)
			return ret
		}
	}
	for i := 0; i < len(in); i++ {
		if in[i].Val != inVal[i] || in[i].isRed != inRed[i] {
			return ret
		}
	}
	return true
}
func TestInsert(t *testing.T) {
	var tree = &RBTree{}
	var vals = []int{8, 18, 5, 15, 17, 25, 40, 80}
	var expPreVal = []int{17, 8, 5, 15, 25, 18, 40, 80}
	var expPreRed = []bool{false, true, false, false, true, false, false, true}
	var expInVal = []int{5, 8, 15, 17, 18, 25, 40, 80}
	var expInRed = []bool{false, true, false, false, false, true, false, true}
	t.Run("Test 1: Insert values", func(t *testing.T) {
		for _, val := range vals {
			tree.Insert(val)
		}
		// tree.PrintPreorder()
		// tree.PrintInorder()
		if !compare(tree, expPreVal, expPreRed, expInVal, expInRed) {
			t.Error("Insert did not work")
		}
	})
	// tree.Insert(5)
	// var a = preorder(tree.Root)
}
func TestSearch(t *testing.T) {
	var tree1 = &RBTree{}
	var tree2 = &RBTree{}
	var vals = []int{8, 18, 5, 15, 17, 25, 40, 80}
	var search1 = 15
	var search2 = 50
	t.Run("Test 1: Check existing value", func(t *testing.T) {
		for _, val := range vals {
			tree1.Insert(val)
		}
		if tree1.Search(search1) == nil {
			t.Error("Search for existing value did not work")
		}
	})
	t.Run("Test 2: Check non-existing value", func(t *testing.T) {
		for _, val := range vals {
			tree2.Insert(val)
		}
		if tree2.Search(search2) != nil {
			t.Error("Search for non-existing value did not work")
		}
	})
}

func TestDelete(t *testing.T) {

	var tree1 = &RBTree{}
	var tree2 = &RBTree{}
	var vals = []int{8, 18, 5, 15, 17, 25, 40, 80}
	var delete1 = 50
	var delete2 = 5
	var expPreVal = []int{17, 8, 5, 15, 25, 18, 40, 80}
	var expPreRed = []bool{false, true, false, false, true, false, false, true}
	var expInVal = []int{5, 8, 15, 17, 18, 25, 40, 80}
	var expInRed = []bool{false, true, false, false, false, true, false, true}
	var deleteExpPreVal = []int{17, 8, 15, 25, 18, 40, 80}
	var deleteExpPreRed = []bool{false, true, false, true, false, false, true}
	var deleteExpInVal = []int{8, 15, 17, 18, 25, 40, 80}
	var deleteExpInRed = []bool{true, false, false, false, true, false, true}
	t.Run("Test 1: Delete non-existing item", func(t *testing.T) {

		for _, val := range vals {
			tree1.Insert(val)
		}
		tree1.Delete(delete1)
		if !compare(tree1, expPreVal, expPreRed, expInVal, expInRed) {
			t.Error("Insert did not work")
		}
	})
	t.Run("Test 2: Delete existing item", func(t *testing.T) {
		for _, val := range vals {
			tree2.Insert(val)
		}
		tree2.Delete(delete2)
		if !compare(tree2, deleteExpPreVal, deleteExpPreRed, deleteExpInVal, deleteExpInRed) {
			t.Error("Insert did not work")
		}
	})
}
