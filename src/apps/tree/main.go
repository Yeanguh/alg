package main

import (
	"apps/tree/comm"
	"apps/tree/in"
	"apps/tree/post"
	"apps/tree/pre"
	"fmt"
)


var Tree *comm.TreeNode

/*
					   20
			15				       30
		9		17		    25		        32
	8				19		            31
*/
func InitTree() {
	if Tree == nil {
		Tree = &comm.TreeNode{}
	}
	Tree.Val = 20
	Tree.Left = &comm.TreeNode{Val: 15}
	Tree.Right = &comm.TreeNode{Val: 30}
	Tree.Left.Left = &comm.TreeNode{Val: 9}
	Tree.Left.Right = &comm.TreeNode{Val: 17}
	Tree.Right.Left = &comm.TreeNode{Val: 25}
	Tree.Right.Right = &comm.TreeNode{Val: 32}
	Tree.Left.Left.Left = &comm.TreeNode{Val: 8}
	Tree.Left.Right.Right= &comm.TreeNode{Val: 19}
	Tree.Right.Right.Left = &comm.TreeNode{Val: 31}
}

func main() {
	InitTree()

	fmt.Println("递归：前",pre.TravelByRecurse(Tree))
	fmt.Println("迭代：前",pre.TravelByStack(Tree))

	fmt.Println("递归：中",in.TravelByRecurse(Tree))
	fmt.Println("迭代：中",in.TravelByStack(Tree))

	fmt.Println("递归：后",post.TravelByRecurse(Tree))
	fmt.Println("递归：后",post.TravelByStack(Tree))



}
