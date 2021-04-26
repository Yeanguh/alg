package common_ancestor

import "apps/tree/types"

// 最近公共祖先
func LCAByRecurse(root *types.TreeNode, p *types.TreeNode, q *types.TreeNode) *types.TreeNode {

	if root == nil {
		return nil
	}

	// 当函数第一次执行到这里时，说明root和p或q相等。那说明p或q其实就是root，pq只要有一个值==root，那无论另一个节点在任意位置，root就是公共祖先，所以直接返回root就是公共祖先了
	/* 例如下面树结构，当函数第一次执行到这里时(root==10)时，如果q==10,那么无论p在这棵树的任意位置，那么10都是公共祖先
				10
			6        17
		3      7
	*/

	// 当函数非第一次执行到这里时，返回root，是为了表明p/q其中一个节点在这里，向上层传递，告诉上层函数，我在这我在这的意思
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 当函数可以执行到这里时，说明pq都不是root根节点，需要去左右中寻找pq
	left := LCAByRecurse(root.Left, p, q) // 递归查找，pq是不是在root的左节点上
	right := LCAByRecurse(root.Right, p, q)// 递归查找，pq是否在root的右节点上
	if left != nil && right != nil {
		// 如果两者都不为nil，说明pq在root根结点的左右，那root就是公共祖先
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}
