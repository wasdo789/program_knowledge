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
