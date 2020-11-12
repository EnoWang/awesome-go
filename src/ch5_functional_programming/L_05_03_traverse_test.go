package ch5_functional_programming

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}

func (node *TreeNode) Print() {

	fmt.Println(node.val)
}

func (node *TreeNode) Traverse() {
	node.TraverseFunc(func(node *TreeNode) {
		node.Print()
	})
}
func TestTraverse(t *testing.T) {
	root := &TreeNode{3, nil, nil}

	root.left = &TreeNode{1, nil, nil}
	root.right = &TreeNode{5, nil, nil}

	root.right.left = &TreeNode{4, nil, nil}
	root.left.right = &TreeNode{2, nil, nil}

	root.Traverse()

	nodeCount := 0

	root.TraverseFunc(func(node *TreeNode) {
		nodeCount++
	})

	t.Log("Node Count: ", nodeCount)

}
