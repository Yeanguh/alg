package travel

import (
	"apps/tree/models"
)

// 中序遍历

//递归
func InTravelByRecurse(root *models.TreeNode) []int {
	var res []int
	var inOrder func(*models.TreeNode)
	inOrder = func(node *models.TreeNode) {
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
func InTravelByStack(root *models.TreeNode) []int {
	var res []int
	var stack []*models.TreeNode
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