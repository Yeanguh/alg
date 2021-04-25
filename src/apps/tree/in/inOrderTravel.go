package in

import (
	"apps/tree/types"
)

// 中序遍历

//递归
func TravelByRecurse(root *types.TreeNode) []int {
	var res []int
	var inOrder func(*types.TreeNode)
	inOrder = func(node *types.TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}

// 迭代
func TravelByStack(root *types.TreeNode) []int {
	var res []int
	var stack []*types.TreeNode
	for root != nil || len(stack) >0 {
		for root!= nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return  res
}