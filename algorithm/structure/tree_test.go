package structure

import (
	"container/list"
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

func TestGetTreeMaxWidth(t *testing.T) {
	var r *TreeNode
	t.Log(GetTreeMaxWidth(r))

	r = &TreeNode{
		Val: 1,
	}
	r.Left = &TreeNode{
		Val: 3,
	}
	r.Right = &TreeNode{
		Val: 5,
	}
	r.Left.Left = &TreeNode{
		Val: 6,
	}
	t.Log(GetTreeMaxWidth(r))
}

func TestSerialize(t *testing.T) {
	var r *TreeNode
	var l *list.List
	l = list.New()
	PreOrderSerialize(r, l)
	PrintList(l)

	r = &TreeNode{
		Val: 1,
	}
	r.Left = &TreeNode{
		Val: 3,
	}
	r.Right = &TreeNode{
		Val: 5,
	}
	r.Left.Left = &TreeNode{
		Val: 6,
	}
	l = list.New()
	PreOrderSerialize(r, l)
	PrintList(l)

	fmt.Println("PreOrderDeserialize")
	nr := PreOrderDeserialize(l)
	PreTraversal(nr)

	fmt.Println("LevelOrderSerialize")
	l = LevelOrderSerialize(r)
	PrintList(l)
	fmt.Println("LevelOrderdDeserialize")
	nr = LevelOrderdDeserialize(l)
	nl := LevelOrderSerialize(nr)
	PrintList(nl)
}

func TestPNode(t *testing.T) {
	var h *PNode
	n := GetNextPNode(h)
	t.Log(n)

	n1 := &PNode{Val: 1}
	n2 := &PNode{Val: 2}
	n3 := &PNode{Val: 3}
	n4 := &PNode{Val: 4}
	n5 := &PNode{Val: 5}
	h = n1
	n1.Left, n1.Right = n2, n3
	n2.Parent, n3.Parent = n1, n1
	n2.Left, n2.Right = n4, n5
	n4.Parent, n5.Parent = n2, n2
	n = GetNextPNode(n5)
	fmt.Printf("%+v\n", n)
}
