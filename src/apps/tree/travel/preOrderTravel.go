package travel

import (
	"apps/tree/models"
)

// 前序遍历

// 递归
func PreTravelByRecurse(root *models.TreeNode) []int {
	var res []int
	var preOrder func(*models.TreeNode)
	preOrder = func(node *models.TreeNode) {
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
func PreTravelByStack(root *models.TreeNode) []int {
	var res []int
	var stack []*models.TreeNode
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
