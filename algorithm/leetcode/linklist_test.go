package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	var r *Node = nil
	r = ReverseList(r)
	fmt.Println(r)
	r = &Node{
		Val:  1,
		Next: nil,
	}
	r.Print()
	r = ReverseList(r)
	r.Print()

	r.Next = &Node{
		Val:  2,
		Next: nil,
	}
	r.Print()
	r = ReverseList(r)
	r.Print()

	r.Next.Next = &Node{
		Val:  3,
		Next: nil,
	}
	r.Print()
	r = ReverseList(r)
	r.Print()
}

func TestReverseDoubleList(t *testing.T) {
	var r *DoubleNode = nil
	r = ReverseDoubleList(r)
	fmt.Println(r)
	r = &DoubleNode{
		Val:  1,
		Next: nil,
		Pre:  nil,
	}
	r.Print()
	r = ReverseDoubleList(r)
	r.Print()
	r.AddNode(&DoubleNode{
		Val:  2,
		Next: nil,
		Pre:  nil,
	})
	r.Print()
	r = ReverseDoubleList(r)
	r.Print()
	r.AddNode(&DoubleNode{
		Val:  3,
		Next: nil,
		Pre:  nil,
	})
	r.Print()
	r = ReverseDoubleList(r)
	r.Print()
}

func TestPrintPublicSection(t *testing.T) {
	l := &Node{
		Val: 1,
	}
	l.AddNode(&Node{Val: 3})
	l.AddNode(&Node{Val: 4})
	l.AddNode(&Node{Val: 7})
	l.Print()

	r := &Node{
		Val: 2,
	}
	r.AddNode(&Node{Val: 3})
	r.AddNode(&Node{Val: 7})
	r.Print()

	PrintPublicSection(l, r)
}

func TestIsPalindromeList(t *testing.T) {
	var l *Node
	res := IsPalindromeList(l)
	assert.Equal(t, res, true)
	l = InitList([]int{1})
	l.Print()
	res = IsPalindromeList(l)
	assert.Equal(t, res, true)
	l.Print()

	l = InitList([]int{1, 2})
	l.Print()
	res = IsPalindromeList(l)
	assert.Equal(t, res, false)
	l.Print()

	l = InitList([]int{1, 2, 1})
	l.Print()
	res = IsPalindromeList(l)
	assert.Equal(t, res, true)
	l.Print()

	l = InitList([]int{1, 2, 2, 1})
	l.Print()
	res = IsPalindromeList(l)
	assert.Equal(t, res, true)
	l.Print()
}

func TestPartition3List(t *testing.T) {
	var l *Node
	l = Partition3List(l, 0)
	fmt.Println(l)

	l = InitList([]int{1})
	l.Print()
	l = Partition3List(l, 1)
	l.Print()

	l = InitList([]int{3, 2})
	l.Print()
	l = Partition3List(l, 2)
	l.Print()

	l = InitList([]int{3, 2, 4, 1})
	l.Print()
	l = Partition3List(l, 2)
	l.Print()

}

func TestCopyRandList(t *testing.T) {
	var r *RandNode
	cr := CopyRandList(r)
	fmt.Println(cr)

	r = &RandNode{
		Val: 1,
	}
	cr = CopyRandList(r)
	fmt.Println(r)
	fmt.Println(cr)

	r.Next = &RandNode{
		Val:  2,
		Rand: r,
	}
	r.Print()
	cr = CopyRandList(r)
	r.Print()
	cr.Print()

	r.Next.Next = &RandNode{
		Val:  3,
		Rand: r.Next,
	}

	r.Print()
	cr = CopyRandList(r)
	r.Print()
	cr.Print()
}

func TestCheckListCycle(t *testing.T) {
	var r *Node
	e := CheckListCycle(r)
	fmt.Println(e)
	r = &Node{
		Val: 1,
	}
	e = CheckListCycle(r)
	fmt.Println(e)

	r.AddNode(&Node{
		Val: 2,
	})
	r.Print()
	e = CheckListCycle(r)
	fmt.Println(e)

	r.AddNode(&Node{
		Val: 3,
	})
	r.Print()
	r.Next.Next.Next = r.Next
	e = CheckListCycle(r)
	fmt.Println(e)
}

func TestCheckListMergeNode(t *testing.T) {
	var l, r *Node
	m := CheckListMergeNode(l, r)
	PrintNode(m)

	l = &Node{
		Val: 1,
	}
	r = &Node{
		Val: 2,
	}
	m = CheckListMergeNode(l, r)
	PrintNode(m)

	l.AddNode(&Node{
		Val: 11,
	})
	publicNode := &Node{
		Val: 8,
	}
	l.Next.Next = publicNode
	r.Next = publicNode
	PrintNode(l)
	PrintNode(r)
	m = CheckListMergeNode(l, r)
	PrintNode(m)

}

func TestAbc(t *testing.T) {
	// x := 10
	// p1 := &x
	// p2 := &x
	// fmt.Printf("%p, %p, %v, %v\n", p1, p2, &x, p1 == p2)

	// n := &Node{
	// 	Val: 1,
	// }
	// p3 := &n
	// p4 := &n
	// fmt.Printf("%p, %p, %p, %v\n", &p3, p3, p4, &n)

	// h := NewList([]int{3, 2, 1})
	// h.Print()

	// r := reverseBetween(h, 1, 2)
	// r.Print()

	h := NewList([]int{1, 2, 3, 4, 5, 6})
	h2 := reverseKGroup(h, 2)
	h2.Print()

}
