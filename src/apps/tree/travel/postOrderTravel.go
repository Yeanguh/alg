package travel

import (
	"apps/tree/models"
)

// 后序遍历

// 递归
func PostTravelByRecurse(root *models.TreeNode) []int {
	var res []int
	var postOrder func(node *models.TreeNode)
	postOrder = func(node *models.TreeNode) {
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
func PostTravelByStack(root *models.TreeNode) []int {
	var res []int
	var stack []*models.TreeNode
	var preN *models.TreeNode
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
