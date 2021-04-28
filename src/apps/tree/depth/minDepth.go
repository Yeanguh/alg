package depth

import (
	"apps/tree/models"
)

func MinDepth(root *models.TreeNode) int {
	var res int
	if root == nil {
		return res
	}
	var stack []*models.TreeNode
	stack = append(stack,root)
	var isRes bool
	for len(stack)>0{
		lens := len(stack)
		for i:=0;i<=lens-1;i++{
			node := stack[0]
			if node.Left != nil{
				stack =append(stack,node.Left)
			}
			if node.Right != nil{
				stack =append(stack,node.Right)
			}
			if node.Left == nil && node.Right == nil {
				isRes = true
			}
		}
		if isRes {
			break
		}
		stack = stack[lens:]
		res++
	}
	return res
}
