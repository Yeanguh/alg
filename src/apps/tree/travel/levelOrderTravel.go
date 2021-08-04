package travel

import (
	"apps/tree/models"
)

// BFS 层序遍历
func LevelTravelByStack(root *models.TreeNode) []int {
	var res []int
	var stack []*models.TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {

		lens := len(stack)
		for i := 0; i <= lens-1; i++ {
			tmp := stack[i]
			res = append(res, tmp.Val)
			if tmp.Left != nil {
				stack = append(stack, tmp.Left)
			}
			if tmp.Right != nil {
				stack = append(stack, tmp.Right)
			}
		}
		stack = stack[lens:]
	}
	return res
}


func ZigzagLevelOrder(root *models.TreeNode) [][]int {
	var flag bool
	var res [][]int
	var stack []*models.TreeNode
	stack = append(stack,root)

	for len(stack)>0 {
		lens := len(stack)
		var tmpRes []int
		for i:=0;i<lens;i++ {
			var node *models.TreeNode
			if !flag {
				node = stack[len(stack)-1]
			} else {
				node = stack[0]
			}
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				stack = append(stack,node.Left)
			}
			if node.Right != nil {
				stack = append(stack,node.Right)
			}
		}
		res = append(res, tmpRes)
		if !flag {
			stack = stack[:lens]
			flag = true
		} else {
			stack = stack[:len(stack)-lens]
			flag =false
		}
	}
	return res
}
