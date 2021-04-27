package depth

import (
	"apps/tree/types"
)

// 最大深度

// DFS深度优先
func MaxDepthByRecurse(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaxDepthByRecurse(root.Left), MaxDepthByRecurse(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 广度优先
func MaxDepthByStack(root *types.TreeNode) int {
	var res int
	var stack []*types.TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		lens := len(stack)
		for i := 0; i <= lens-1; i++ {
			tmp := stack[0]
			if tmp.Left != nil {
				stack = append(stack, tmp.Left)
			}
			if tmp.Right != nil {
				stack = append(stack, tmp.Right)
			}
		}
		stack = stack[lens:]
		res++
	}
	return res
}