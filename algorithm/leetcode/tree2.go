package leetcode

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeExtra struct {
	Node   *TreeNode
	Length uint64
	Height int
}

func TreeDfs(head *TreeNode, m map[int]NodeExtra, height int, length uint64) {
	if head == nil {
		return
	}

	if _, ok := m[height]; !ok {
		m[height] = NodeExtra{
			Node:   head,
			Length: length,
			Height: height,
		}
	}
	if head.Right != nil {
		TreeDfs(head.Right, m, height+1, length*2+1)
	}

	if head.Left != nil {
		TreeDfs(head.Left, m, height+1, length*2)
	}

}

func rightSideView(root *TreeNode) []int {
	var nodes []NodeExtra
	heightNodeMap := map[int]NodeExtra{}
	TreeDfs(root, heightNodeMap, 1, 1)
	for key := range heightNodeMap {
		val := heightNodeMap[key]
		nodes = append(nodes, val)
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Height < nodes[j].Height
	})
	var res []int
	for _, n := range nodes {
		res = append(res, n.Node.Val)
	}
	return res
}
func isSymmetricSub(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && isSymmetricSub(t1.Left, t2.Right) && isSymmetricSub(t1.Right, t2.Left)
	}
	return false
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricSub(root.Left, root.Right)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	if root.Left != nil && hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeRe(&preorder, inorder)
}

func buildTreeRe(preorder *([]int), inorder []int) *TreeNode {
	if len(*preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var head *TreeNode
	for i := 0; i < len(inorder); i++ {
		if (inorder)[i] == (*preorder)[0] {
			head = &TreeNode{
				Val: (*preorder)[0],
			}
			*preorder = (*preorder)[1:]
			if i-1 >= 0 && len(*preorder) > 0 {
				head.Left = buildTreeRe(preorder, inorder[0:i])
			}

			if len(*preorder) > 0 && i+1 < len(inorder) {
				head.Right = buildTreeRe(preorder, inorder[i+1:])
			}
			break
		}
	}
	return head
}

func buildTreeRe2(inorder []int, is int, ie int, postorder []int, ps int, pe int) *TreeNode {

	if is == ie {
		return &TreeNode{
			Val: inorder[is],
		}
	}

	parent := &TreeNode{
		Val: postorder[pe],
	}
	idx := is
	for ; idx <= ie; idx++ {
		if inorder[idx] == parent.Val {
			break
		}
	}
	rightChildsNum := ie - idx
	if idx+1 < len(inorder) && idx+1 <= ie {
		parent.Right = buildTreeRe2(inorder, idx+1, ie, postorder, pe-rightChildsNum, pe-1)
	}
	if idx-1 >= 0 && is <= idx-1 {
		parent.Left = buildTreeRe2(inorder, is, idx-1, postorder, ps, pe-rightChildsNum-1)
	}
	return parent
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	return buildTreeRe2(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//返回头，尾部指针
func flattenre(root *TreeNode) (*TreeNode, *TreeNode) {
	//边界条件TODO
	if root == nil {
		return nil, nil
	}

	head := root
	tail := root
	rightChild := root.Right
	if root.Left != nil {
		h, t := flattenre(root.Left)
		if h != nil {
			root.Right = h
			tail = t
		}
	}
	root.Left = nil
	if rightChild != nil {
		h, t := flattenre(rightChild)
		if h != nil {
			tail.Right = h
			tail = t
		}
	}
	return head, tail
}
func flatten(root *TreeNode) {
	flattenre(root)
}

func dfssumNumbers(presum int, root *TreeNode, sum *int) {
	presum = presum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += presum
	}
	if root.Left != nil {
		dfssumNumbers(presum, root.Left, sum)
	}
	if root.Right != nil {
		dfssumNumbers(presum, root.Right, sum)
	}
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	dfssumNumbers(0, root, &sum)
	return sum
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 以root 为端点的最大路径和
func maxchildpath(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	leftgain := Max(maxchildpath(root.Left, sum), 0)
	rightgain := Max(maxchildpath(root.Right, sum), 0)
	newsum := root.Val + leftgain + rightgain
	if newsum > *sum {
		*sum = newsum
	}
	return Max(root.Val, root.Val+Max(leftgain, rightgain))

}
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val
	maxchildpath(root, &sum)
	return sum

}
