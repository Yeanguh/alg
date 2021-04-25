package post

import (
	"apps/tree/types"
)

// 后序遍历

// 递归
func TravelByRecurse(root *types.TreeNode) []int {
	var res []int
	var postOrder func(node *types.TreeNode)
	postOrder = func(node *types.TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		res = append(res, node.Val)
	}
	postOrder(root)
	return res
}

// 迭代
func TravelByStack(root *types.TreeNode) []int {
	var res []int
	var stack []*types.TreeNode
	var preN *types.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == preN {
			res = append(res, root.Val)
			preN = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
