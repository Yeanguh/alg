package main

import (
	"apps/tree/common_ancestor"
	"fmt"

	"apps/tree/depth"
	"apps/tree/travel"
	"apps/tree/types"
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
	fmt.Println("递归：前",travel.PreTravelByRecurse(Tree))
	fmt.Println("迭代：前",travel.PreTravelByStack(Tree))
	fmt.Println("递归：中", travel.InTravelByRecurse(Tree))
	fmt.Println("迭代：中", travel.InTravelByStack(Tree))
	fmt.Println("递归：后",travel.PostTravelByRecurse(Tree))
	fmt.Println("迭代：后",travel.PostTravelByStack(Tree))
	fmt.Println("迭代：层次", travel.LevelTravelByStack(Tree))

	// 二叉树最大深度
	fmt.Println("递归：最大深度", depth.MaxDepthByRecurse(Tree))
	fmt.Println("迭代：最大深度", depth.MaxDepthByStack(Tree))

	// 最近公共祖先
	ca := common_ancestor.LCAByRecurse(Tree,&types.TreeNode{8,nil,nil},&types.TreeNode{9,nil,nil}).Val
	fmt.Println("最近公共祖先：",ca)

}

