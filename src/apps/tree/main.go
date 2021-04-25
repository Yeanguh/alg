package main

import (
	"apps/tree/lca"
	"apps/tree/types"
	"apps/tree/dep"
	"apps/tree/in"
	"apps/tree/level"
	"apps/tree/post"
	"apps/tree/pre"
	"fmt"
)


var Tree *types.TreeNode

/*
					   20
			15				       30
		9		17		    25		        32
	8				19		            31
*/
func InitTree() {
	if Tree == nil {
		Tree = &types.TreeNode{}
	}
	Tree.Val = 20
	Tree.Left = &types.TreeNode{Val: 15}
	Tree.Right = &types.TreeNode{Val: 30}
	Tree.Left.Left = &types.TreeNode{Val: 9}
	Tree.Left.Right = &types.TreeNode{Val: 17}
	Tree.Right.Left = &types.TreeNode{Val: 25}
	Tree.Right.Right = &types.TreeNode{Val: 32}
	Tree.Left.Left.Left = &types.TreeNode{Val: 8}
	Tree.Left.Right.Right= &types.TreeNode{Val: 19}
	Tree.Right.Right.Left = &types.TreeNode{Val: 31}
}

func main() {
	InitTree()

	// 遍历
	fmt.Println("递归：前",pre.TravelByRecurse(Tree))
	fmt.Println("迭代：前",pre.TravelByStack(Tree))
	fmt.Println("递归：中",in.TravelByRecurse(Tree))
	fmt.Println("迭代：中",in.TravelByStack(Tree))
	fmt.Println("递归：后",post.TravelByRecurse(Tree))
	fmt.Println("迭代：后",post.TravelByStack(Tree))
	fmt.Println("迭代：层次",level.TravelByStack(Tree))

	// 二叉树最大深度
	fmt.Println("递归：最大深度", dep.MaxDepth(Tree))
	fmt.Println("迭代：最大深度", level.MaxDepth(Tree))

	// 最近公共祖先
	fmt.Println(lca.LCA(Tree,&types.TreeNode{8,nil,nil},&types.TreeNode{9,nil,nil}).Val)

}

