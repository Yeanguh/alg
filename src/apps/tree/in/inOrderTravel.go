package in

import (
	"apps/tree/comm"
)

// 中序遍历

//递归
func TravelByRecurse(root *comm.TreeNode) []int {
	var res []int
	var inOrder func(*comm.TreeNode)
	inOrder = func(node *comm.TreeNode) {
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
func TravelByStack(root *comm.TreeNode) []int {
	var res []int
	var stack []*comm.TreeNode
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