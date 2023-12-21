package structure

import (
	"fmt"
	"testing"
)

func TestTreeNode(t *testing.T) {

	r := &TreeNode{
		Val: 1,
	}
	r.Left = &TreeNode{
		Val: 3,
	}
	r.Right = &TreeNode{
		Val: 5,
	}
	PreTraversal(nil)
	PreTraversal(r)
	PreTraversalNonRecursion(r)
	fmt.Println("in order")
	InTraversal(nil)
	InTraversal(r)
	InTraversalNonRecursion(r)

	fmt.Println("post order")
	PostTraversal(nil)
	PostTraversal(r)
	PostTraversalNonRecursion(r)
}
