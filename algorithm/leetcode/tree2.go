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
