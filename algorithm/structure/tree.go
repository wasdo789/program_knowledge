package structure

import (
	"container/heap"
	"container/list"
	"fmt"
	"strconv"
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

// 前序的序列化
func PreOrderSerialize(r *TreeNode, res *list.List) {
	if r == nil {
		res.PushBack("null")
		return
	} else {
		res.PushBack(fmt.Sprint(r.Val))
	}

	PreOrderSerialize(r.Left, res)
	PreOrderSerialize(r.Right, res)

}

// 根据前序遍历数组构造二叉树，中左右
func PreOrderDeserialize(res *list.List) *TreeNode {
	//
	v := res.Remove(res.Front()).(string)
	if v == "null" {
		return nil
	} else {
		v, _ := strconv.ParseInt(v, 10, 64)
		r := &TreeNode{
			Val: int(v),
		}
		r.Left = PreOrderDeserialize(res)
		r.Right = PreOrderDeserialize(res)
		return r
	}

}

func PrintList(l *list.List) {
	for r := l.Front(); r != nil; r = r.Next() {
		fmt.Println(r.Value)
	}
}

// 获取二叉树最大宽度和层数
func GetTreeMaxWidth(r *TreeNode) (int, int) {
	if r == nil {
		return 0, 0
	}
	curLevelWidth := 1
	curLevel := 1
	nextLevelWidth := 0

	maxWidth := 1
	maxWidthLevel := 1
	queue := list.New()
	queue.PushBack(r)
	for queue.Len() > 0 {
		e := queue.Remove(queue.Front()).(*TreeNode)
		curLevelWidth--
		if e.Left != nil {
			queue.PushBack(e.Left)
			nextLevelWidth++
		}
		if e.Right != nil {
			queue.PushBack(e.Right)
			nextLevelWidth++
		}

		if curLevelWidth == 0 {
			if maxWidth < nextLevelWidth {
				maxWidth = nextLevelWidth
				maxWidthLevel = curLevel + 1
			}
			curLevel++
			curLevelWidth = nextLevelWidth
			nextLevelWidth = 0
		}
	}
	return maxWidthLevel, maxWidth
}

// 层序序列化
func LevelOrderSerialize(r *TreeNode) *list.List {
	l := list.New()
	if r == nil {
		return l
	}
	queue := list.New()
	queue.PushBack(r)
	for queue.Len() > 0 {
		e := queue.Remove(queue.Front()).(*TreeNode)
		if e == nil {
			l.PushBack("null")
		} else {
			l.PushBack(fmt.Sprint(e.Val))
			queue.PushBack(e.Left)
			queue.PushBack(e.Right)
		}
	}
	return l
}

func NewNodeByVal(val string) *TreeNode {
	if val == "null" {
		return nil
	} else {
		v, _ := strconv.ParseInt(val, 10, 64)
		return &TreeNode{
			Val: int(v),
		}
	}
}

// 层序反序列化
func LevelOrderdDeserialize(l *list.List) *TreeNode {
	if l.Len() == 0 {
		return nil
	}
	queue := list.New()
	v := l.Remove(l.Front()).(string)
	root := NewNodeByVal(v)
	if root == nil {
		return nil
	}
	queue.PushBack(root)
	for queue.Len() > 0 && l.Len() > 0 {
		curParent := queue.Remove(queue.Front()).(*TreeNode)
		v := l.Remove(l.Front()).(string)
		curNode := NewNodeByVal(v)
		curParent.Left = curNode
		if curNode != nil {
			queue.PushBack(curNode)
		}
		if l.Len() > 0 {
			v := l.Remove(l.Front()).(string)
			curNode := NewNodeByVal(v)
			curParent.Right = curNode
			if curNode != nil {
				queue.PushBack(curNode)
			}
		}
	}
	return root
}

// 有parent指针的节点
type PNode struct {
	Val    int
	Left   *PNode
	Right  *PNode
	Parent *PNode
}

// 找某个节点的中序遍历的后继节点
func GetNextPNode(n *PNode) *PNode {
	if n == nil {
		return nil
	}
	//有右节点,找到右子树的最左节点
	if n.Right != nil {
		cur := n.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	} else {
		// 无右子树，找
		parent := n.Parent
		for n != parent.Left && parent != nil {
			n = parent
			parent = parent.Parent
		}
		return parent
	}
}

//

/*
哈夫曼树定义：树的带权路径长度（Weighted Path Length of Tree，WPL） 最小的二叉树 称为 霍夫曼树（Huffman Tree）
哈夫曼树构造：每次从小顶堆中取得最小的两个
*/
func Huffman(datas []int) *TreeNode {
	if len(datas) == 0 {
		return nil
	}
	var hp TreeNodeHeap
	for _, d := range datas {
		hp = append(hp, &TreeNode{Val: d})
	}
	//有限队列
	heap.Init(&hp)
	for len(hp) > 0 {
		x := heap.Pop(&hp).(*TreeNode)
		if len(hp) == 0 {
			return x
		}
		y := heap.Pop(&hp).(*TreeNode)
		merged := &TreeNode{
			Val: x.Val + y.Val,
		}
		merged.Left = x
		merged.Right = y
		heap.Push(&hp, merged)

	}
	return nil
}
