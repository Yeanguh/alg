package main

import (
	"apps/linklist/models"
	"apps/linklist/single"
	"fmt"
)

var LinkList *models.ListNode

func init(){
	LinkList = &models.ListNode{Val:1,Next:&models.ListNode{Val:2,Next:&models.ListNode{Val:3,Next:&models.ListNode{Val:4,Next:&models.ListNode{Val:5,Next:nil}}}}}
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
	printLinklist(single.Reverse(LinkList))
}