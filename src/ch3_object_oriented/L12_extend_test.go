package ch3_object_oriented

import "testing"

type MyTreeNode struct {
	node *TreeNode
}

// 0. Extend a struct by combination
// Note: can not write in MyTreeNode{myNode.node.left}.postOrder() bc we can only get pointer for a variable (can not literal)
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.left}
	left.postOrder()
	right := MyTreeNode{myNode.node.right}
	right.postOrder()
	myNode.node.print()
}

func TestPostOrder(t *testing.T) {

	root := &TreeNode{3, nil, nil}

	root.left = &TreeNode{1, nil, nil}
	root.right = &TreeNode{5, nil, nil}

	root.right.left = &TreeNode{4, nil, nil}
	root.left.right = &TreeNode{2, nil, nil}

	myNode := &MyTreeNode{root}
	myNode.postOrder()
}

// 1. Extend a struct by alias

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Poll() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func TestQueue(t *testing.T) {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	t.Log(q.Poll())
	t.Log(q.Poll())
	t.Log(q.IsEmpty())
	t.Log(q.Poll())
	t.Log(q.IsEmpty())
}
