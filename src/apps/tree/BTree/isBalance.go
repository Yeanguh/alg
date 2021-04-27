package BTree

import (
	"apps/tree/types"
)

func IsBalancedByRecurse(root *types.TreeNode) bool {
	return pro(root) != -1
}

func pro(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := pro(root.Left), pro(root.Right)
	if l == -1 || r == -1 {
		return -1
	}

	if l - r > 1 || r - l > 1 {
		return -1
	}

	return m(l, r) + 1
}

func m(a,b int) int {
	if a > b {
		return a
	}

	return b
}