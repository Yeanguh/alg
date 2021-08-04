package main

import (
	"apps/linklist/models"
	"apps/linklist/single"
	"fmt"
)

var xLinkList *models.ListNode
var singleLinkList *models.ListNode

func init(){
	node := &models.ListNode{Val:3,Next:nil}

	xLinkList = &models.ListNode{Val:1, Next:&models.ListNode{Val: 2, Next:node}}
	node.Next = &models.ListNode{Val:4,Next:&models.ListNode{Val:5,Next:node}}

	singleLinkList = &models.ListNode{Val:1,Next:&models.ListNode{Val:2,Next:&models.ListNode{Val:3,Next:&models.ListNode{Val:4,Next:nil}}}}
}

func printLinklist(head *models.ListNode){
	fmt.Println(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("----分界线-----")
}

func main(){
	// 是否有环
	fmt.Println(single.HasRingByMap(xLinkList))
	fmt.Println(single.HasRingByFastSlowPointer(xLinkList))
	fmt.Println(single.HasRingByFastSlowPointer(singleLinkList))

	// 反转链表
	printLinklist(single.Reverse(singleLinkList))

	printLinklist(single.ReverseKGroup1(singleLinkList,3))

}