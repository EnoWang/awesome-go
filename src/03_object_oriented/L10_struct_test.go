package _3_object_oriented

import (
	"fmt"
	"testing"
)

// 0. There is only Encapsulation in golang, no Inheritance or Polymorphism
// Note: There is no class in golang
// Note: no need to know whether struct create in heap or stack

// 1. create struct
// Note: always use . to dereference struct, no matter pointer or instance
type TreeNode struct {
	val         int
	left, right *TreeNode
}

func TestBasicStruct(t *testing.T) {
	var root TreeNode
	t.Log(root)

	// solution 1
	root = TreeNode{
		val: 3,
	}

	// solution 2
	root.left = &TreeNode{4, nil, nil}
	root.right = &TreeNode{5, nil, nil}

	// solution 3
	root.right.left = new(TreeNode)
}

// 2. create struct slice
func TestStructSlice(t *testing.T) {

	// Note: create a slice
	nodes := []TreeNode{
		{val: 3},
		{},
		{6, nil, nil},
	}

	t.Log(nodes)
}

// 3. factory function
// Note: no construct function in golang, can add factory function
func createNode(val int) *TreeNode {
	// Note: return the address for a local variable
	// No work in c++, but work in golang
	return &TreeNode{val: val}
}

func TestFactory(t *testing.T) {

	node := createNode(100)
	t.Log(*node)
}

// 4. create method for struct
/* Note: method in golang is the same as a general function
we can also write this method as : func print(node TreeNode) and call it as print(node)
So node is pass by value instead of reference, so setValue doesn't work.
That's why we always add method on pointer as below
*/

func (node TreeNode) print() {

	fmt.Println(node.val)
}

func (node TreeNode) setValue(val int) {
	node.val = val
}

func TestMethod(t *testing.T) {

	nodePtr := createNode(100)
	node := *nodePtr
	node.print()
	node.setValue(10)
	node.print()

	nodePtr.print()
	nodePtr.setValue(10)
	nodePtr.print()
}

// 5. create method for struct pointer
func (node *TreeNode) setNodeValue(val int) {
	node.val = val
}

// Note: we can call setNodeValue by pointer or struct itself
// regards on the receiver of method, compiler will copy the value (if TreeNode) or pass the pointer (if *TreeNode)
func TestMethodOnPointer(t *testing.T) {

	node := createNode(100)
	node.print()
	node.setNodeValue(20)
	node.print()
	(*node).setNodeValue(30)
	node.print()
}

// 5. traverse tree
// Note: nil could be pass as pointer safely
func (node *TreeNode) traverseTree() {
	if node == nil {
		return
	}
	node.left.traverseTree()
	node.print()
	node.right.traverseTree()
}

func TestTraverseTree(t *testing.T) {

	root := &TreeNode{3, nil, nil}

	root.left = &TreeNode{1, nil, nil}
	root.right = &TreeNode{5, nil, nil}

	root.right.left =  &TreeNode{4, nil, nil}
	root.left.right = &TreeNode{2, nil, nil}

	root.traverseTree()
}


