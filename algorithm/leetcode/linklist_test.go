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
