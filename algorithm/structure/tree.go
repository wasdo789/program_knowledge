package structure

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreTraversal(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)

	if t.Left != nil {
		PreTraversal(t.Left)
	}
	if t.Right != nil {
		PreTraversal(t.Right)
	}
}

func InTraversal(t *TreeNode) {
	if t == nil {
		return
	}
	if t.Left != nil {
		InTraversal(t.Left)
	}
	fmt.Println(t.Val)
	if t.Right != nil {
		InTraversal(t.Right)
	}
}

func PostTraversal(t *TreeNode) {
	if t == nil {
		return
	}
	if t.Left != nil {
		PreTraversal(t.Left)
	}
	if t.Right != nil {
		PreTraversal(t.Right)
	}
	fmt.Println(t.Val)
}

func PreTraversalNonRecursion(r *TreeNode) {
	if r == nil {
		return
	}
	//双向链表模拟栈
	s := list.New()
	s.PushFront(r)
	for s.Len() != 0 {
		e := s.Remove(s.Front()).(*TreeNode)
		fmt.Println(e.Val)
		if e.Right != nil {
			s.PushFront(e.Right)
		}
		if e.Left != nil {
			s.PushFront(e.Left)
		}
	}
}

func InTraversalNonRecursion(t *TreeNode) {
	if t == nil {
		return
	}
	s := list.New()
	// if t.Left != nil {
	// 	PreTraversal(t.Left)
	// }
	// fmt.Println(t.Val)
	// if t.Right != nil {
	// 	PreTraversal(t.Right)
	// }
	for t != nil {
		s.PushFront(t)
		t = t.Left
	}
	for s.Len() != 0 {
		e := s.Remove(s.Front()).(*TreeNode)
		fmt.Println(e.Val)
		if e.Right != nil {
			s.PushFront(e.Right)
		}
	}
}

func PostTraversalNonRecursion(t *TreeNode) {
	if t == nil {
		return
	}
	//左、右、中
	s := list.New()
	// if t.Left != nil {
	// 	PreTraversal(t.Left)
	// }
	// fmt.Println(t.Val)
	// if t.Right != nil {
	// 	PreTraversal(t.Right)
	// }
	lastPop := t
	s.PushFront(t)

	for s.Len() != 0 {
		cur := s.Front().Value.(*TreeNode)
		if cur.Left != nil && lastPop != cur.Left && lastPop != cur.Right {
			s.PushFront(cur.Left)
		} else if cur.Right != nil && lastPop != cur.Right {
			s.PushFront(cur.Right)
		} else {
			fmt.Println(cur.Val)
			e := s.Remove(s.Front()).(*TreeNode)
			lastPop = e
		}
	}
}
