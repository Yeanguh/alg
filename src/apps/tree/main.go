package main

import (
	"apps/tree/BTree"
	"apps/tree/LCA"
	"apps/tree/symmetric"
	"fmt"

	"apps/tree/depth"
	"apps/tree/travel"
	"apps/tree/models"
)


var Tree *models.TreeNode

/*
					   20
			15				       30
		9		17		    25		        32
	8				19		            31
*/
func InitTree() {
	if Tree == nil {
		Tree = &models.TreeNode{}
	}
	Tree.Val = 20
	Tree.Left = &models.TreeNode{Val: 15}
	Tree.Right = &models.TreeNode{Val: 30}
	Tree.Left.Left = &models.TreeNode{Val: 9}
	Tree.Left.Right = &models.TreeNode{Val: 17}
	Tree.Right.Left = &models.TreeNode{Val: 25}
	Tree.Right.Right = &models.TreeNode{Val: 32}
	Tree.Left.Left.Left = &models.TreeNode{Val: 8}
	Tree.Left.Right.Right= &models.TreeNode{Val: 19}
	Tree.Right.Right.Left = &models.TreeNode{Val: 31}
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
	// 二叉树最小深度
	fmt.Println("迭代：最小深度", depth.MinDepth(Tree))

	fmt.Println("是否是对称二叉树：",symmetric.IsSymmetric(Tree))
	fmt.Println("是否是平衡二叉树：",BTree.IsBalancedByRecurse(Tree)) // 不好理解

	fmt.Println("最近公共祖先：",LCA.ByRecurse(Tree,&models.TreeNode{8,nil,nil},&models.TreeNode{9,nil,nil}).Val)
}

