package pre

import (
	"apps/tree/types"
)

// 前序遍历

// 递归
func TravelByRecurse(root *types.TreeNode) []int {
	var res []int
	var preOrder func(*types.TreeNode)
	preOrder = func(node *types.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return res
}

// 迭代
func TravelByStack(root *types.TreeNode) []int {
	var res []int
	var stack []*types.TreeNode
	for root != nil || len(stack)>0 {
		for root !=nil {
			res = append(res, root.Val)
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
