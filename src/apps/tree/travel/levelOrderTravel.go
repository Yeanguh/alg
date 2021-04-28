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
