package symmetric

import "apps/tree/models"

func IsSymmetric(root *models.TreeNode) bool {
	if root == nil {
		return true
	}
	return nodeIsEq(root.Left, root.Right)
}

func nodeIsEq (leftNode, rightNode *models.TreeNode) bool {
	if leftNode == nil && rightNode ==nil {
		return true
	}
	if (leftNode == nil && rightNode !=nil)||(leftNode != nil && rightNode ==nil) || leftNode.Val != rightNode.Val {
		return false
	}
	return nodeIsEq(leftNode.Left, rightNode.Right) && nodeIsEq(leftNode.Right,rightNode.Left)
}
