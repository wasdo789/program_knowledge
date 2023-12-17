package leetcode

/*
	链表相关算法题
*/

import (
	"fmt"
)

// 单链表结构
type Node struct {
	Val  int
	Next *Node
}

func (n *Node) AddNode(i *Node) {
	for ; n.Next != nil; n = n.Next {
	}
	n.Next = i
}

func InitList(vals []int) *Node {
	var head *Node
	if len(vals) >= 1 {
		head = &Node{Val: vals[0]}
	}
	for _, val := range vals[1:] {
		head.AddNode(&Node{
			Val: val,
		})
	}
	return head
}

func (n Node) Print() {
	for cur := &n; cur != nil; cur = cur.Next {
		fmt.Printf("addr:%p, %+v->", cur, cur)
	}
	fmt.Println()
}

func PrintNode(n *Node) {
	if n == nil {
		fmt.Print("nil")
	} else {
		for cur := n; cur != nil; cur = cur.Next {
			fmt.Printf("addr:%p, %+v->", cur, cur)
		}
	}
	fmt.Println()
}

// 双链表结构
type DoubleNode struct {
	Val  int
	Next *DoubleNode
	Pre  *DoubleNode
}

func (n *DoubleNode) AddNode(i *DoubleNode) {
	for ; n.Next != nil; n = n.Next {
	}
	n.Next = i
	if i != nil {
		i.Pre = n
	}
}

func (n DoubleNode) Print() {
	for cur := &n; cur != nil; cur = cur.Next {
		fmt.Printf("%+v->", cur)
	}
	fmt.Println()
}

// 反转单链表
func ReverseList(root *Node) *Node {
	if root == nil || root.Next == nil {
		return root
	}
	next := root.Next
	root.Next = nil
	for next != nil {
		afterNext := next.Next
		next.Next = root
		root = next
		next = afterNext
	}
	return root
}

// 反转双链表
func ReverseDoubleList(root *DoubleNode) *DoubleNode {
	if root == nil || root.Next == nil {
		return root
	}
	for {
		tmp := root.Next
		root.Next, root.Pre = root.Pre, root.Next
		if tmp == nil {
			break
		}
		root = tmp
	}
	return root
}

// 打印两个有序链表的公共部分
// 比如1->3->4->7   2->3->7  打印3->7
func PrintPublicSection(l *Node, r *Node) {
	for l != nil && r != nil {
		if l.Val == r.Val {
			fmt.Println(l.Val)
			l = l.Next
			r = r.Next
		} else if l.Val < r.Val {
			l = l.Next
		} else {
			r = r.Next
		}
	}
}

// 判断链表是否为回文结构，即中心对称
// 1->1 是，1->2->1是，1->2->2->1是，1->2-3不是
// 简单方案1，使用栈辅助，额外空间复杂度为o(n)，太简单不实现了
// 复杂方案2，使用快慢指针找到中间节点，然后逆转前半部分链表，然后从中间使用双指针对比
func IsPalindromeList(l *Node) bool {
	if l == nil || l.Next == nil {
		return true
	}
	if l.Next.Next == nil {
		return l.Val == l.Next.Val
	}
	fast, slow := l, l
	var preSlow *Node //slow前一个
	//fast一次两步，slow一次一步，fast到结尾的时候，slow期望指向中间
	//同时逆转链表前半部分，直到slwo
	//截止条件1，fast->next == nil（奇数），slow指向中间，fast指向末尾
	//截止条件2， fast->next-Next == nil（偶数），slow指向前半部分最后一个，fast指向倒数第二个
	for fast.Next != nil && fast.Next.Next != nil {
		//逆转slow之前的节点
		tmp := slow.Next //暂存下个位置
		fast = fast.Next.Next

		//逆转slow之前的
		slow.Next = preSlow
		//向后移动
		preSlow = slow
		slow = tmp

	}
	//首先保存slow的下一个位置，用于恢复
	nextSlow := slow.Next
	var left, right *Node
	// 对比，奇数，如果fast.Next == nil 那么左边右边开始位置：slow 开始对比
	if fast.Next == nil {
		slow.Next = preSlow
		left, right = slow, slow
	} else if fast.Next.Next == nil {
		// 偶数，若果fast->next-Next == nil，右边开始位置slow->next，左边开始位置slow
		right = slow.Next
		slow.Next = preSlow
		left = slow
	}
	res := true
	for left != nil && right != nil {
		if left.Val != right.Val {
			res = false
			break
		}
		left = left.Next
		right = right.Next
	}
	//恢复链表
	for slow != nil {
		//保持下个位置
		tmp := slow.Next
		//反转
		slow.Next = nextSlow
		//更新位置
		nextSlow = slow
		slow = tmp
	}
	return res
}

// 三分链表
// 将单向链表按照某值划分成左边小，中间等，右边大的
// 方法1，把链表节点放入数组进行分区
// 方法2，直接分区
func Partition3List(r *Node, val int) *Node {
	//维护6个指针，表示小于区开头结尾的，sh，st
	//等于区范围的eh，et，大于区的bh，bt
	//依次扫描，遇到小于的，插入st后面，遇到等于的，插入et后面，遇到大于的
	//插入bt后面
	//最后把st->eh，et->bh 连接起来即可
	if r == nil || r.Next == nil {
		return r
	}
	var sh, st, eh, et, bh, bt *Node
	cur := r
	for cur != nil {
		tmp := cur.Next
		cur.Next = nil
		if cur.Val < val {
			if st == nil {
				sh, st = cur, cur
			} else {
				st.Next = cur
				st = cur
			}
		} else if cur.Val == val {
			if et == nil {
				eh, et = cur, cur
			} else {
				et.Next = cur
			}
		} else {
			if bt == nil {
				bh, bt = cur, cur
			} else {
				bt.Next = cur
			}
		}
		cur = tmp
	}

	if st != nil {
		if et != nil {
			st.Next = eh
			et.Next = bh
		} else {
			st.Next = bh
		}
		return sh
	} else {
		if et != nil {
			et.Next = bh
			return eh
		} else {
			return bh
		}
	}

}

type RandNode struct {
	Val  int
	Next *RandNode
	Rand *RandNode
}

func (n RandNode) Print() {
	for cur := &n; cur != nil; cur = cur.Next {
		if cur.Rand != nil {
			fmt.Printf("val:%v, rand:%v", cur.Val, cur.Rand.Val)
		} else {
			fmt.Printf("val:%v, rand:%v", cur.Val, cur.Rand)
		}

	}
	fmt.Println()
}

// 复制特殊链表，该链表有个额外的随机指针
// 简单方法1，使用has结构存储旧节点到新节点的映射，不实现
// 复杂方法2，依次在原链表每个节点后面新增一个复制节点，最后把
// 复制节点连接起来即可。复制节点的随机指针指向原节点的随机指针的下一个
func CopyRandList(r *RandNode) *RandNode {
	if r == nil {
		return nil
	} else if r.Next == nil {
		return &RandNode{
			Val: r.Val,
		}
	}

	//复制节点
	var h, cur *RandNode
	for cur = r; cur != nil; {
		next := cur.Next
		cpNode := &RandNode{
			Val: cur.Val,
		}
		if h == nil {
			h = cpNode
		}
		cur.Next = cpNode
		cpNode.Next = next
		cur = next
	}
	//复制随机指针,cur依次往后移动，指向链表的旧节点
	for cur = r; cur != nil; {
		next := cur.Next.Next
		if cur.Rand != nil {
			cur.Next.Rand = cur.Rand.Next
		}
		cur = next

	}
	//恢复链表
	cur = r
	var preNewNode *RandNode
	for cur != nil {
		nextOldNode := cur.Next.Next //先保存下个旧链表节点位置
		if preNewNode != nil {
			preNewNode.Next = cur.Next
		}
		preNewNode = cur.Next
		cur.Next = nextOldNode
		cur = nextOldNode
	}
	return h

}

// 检测链表是否有环，并且返回环的入口
// 简单方法，使用辅助hash结构存储地址，如果遇到相同的，说明有环
// 使用快慢指针，如果有环，快指针会追上慢指针，然后把快指针指向链表头，
// 快慢指针每次走一步，直至相遇，就是环的入口
func CheckListCycle(r *Node) *Node {
	if r == nil || r.Next == nil {
		return nil
	}
	slow := r.Next
	fast := r.Next.Next
	//如果快指针变为nil，说明无环
	for slow != fast {
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
	}
	fast = r
	for fast = r; fast != slow; {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

/*
检测两个无环链表是否相交，并返回首个相交节点
如果相交，两个链表末尾节点一定相等
假设两个链表长度相差a，长指针先走n，然后短链表指针同时移动，节点相等时就是首个相交节点
*/
func CheckListMergeNode(l *Node, r *Node) *Node {
	if l == nil || r == nil {
		return nil
	}
	lCur, rCur := l, r
	var lSize, rSize int
	for ; lCur.Next != nil; lCur = lCur.Next {
		lSize++
	}
	for ; rCur.Next != nil; rCur = rCur.Next {
		rSize++
	}
	if lCur != rCur {
		return nil
	}
	lCur, rCur = l, r
	if rSize > lSize {
		for i := 0; i < rSize-lSize; i++ {
			rCur = rCur.Next
		}
	} else {
		for i := 0; i < lSize-rSize; i++ {
			lCur = lCur.Next
		}
	}
	for lCur != rCur {
		lCur = lCur.Next
		rCur = rCur.Next
	}

	return lCur
}
