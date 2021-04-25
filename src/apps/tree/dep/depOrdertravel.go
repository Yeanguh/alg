package dep

import (
	"apps/tree/types"
)

// 最大深度

// DFS 深度优先遍历

func MaxDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}