package single

import "apps/linklist/models"

// 链表反转
func Reverse(head *models.ListNode) *models.ListNode {

	if head == nil {
		return nil
	}
	var newHead *models.ListNode // 声明一个新的头节点
	for head != nil {
		tmp := head.Next // 先把下一个节点临时储存起来
		head.Next = newHead // 将当前节点next指针指向nil（因为新节点newHead当前为nil），该节点处于孤立，没有节点相连
		newHead = head // 然后让新节点等于当前节点
		head = tmp // 再把原节点head移动到下一个节点上，因为已经没有指针指向，直接把开始储存tmp赋值给head，相当于移动操作
	}
	return newHead
}

// https://leetcode-cn.com/problems/palindrome-linked-list/
func IsPalindrome(head *models.ListNode) bool {
	return false
}
func ReverseKGroup1(head *models.ListNode, k int) *models.ListNode {
	dummy := &models.ListNode{}
	dummy.Next = head
	pre := dummy
	end := dummy
	for end.Next != nil {
		for i:=0;i<k&&end!=nil;i++ {
			end = end.Next
		}
		if end == nil {
			return dummy.Next
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = Reverse(start)

		start.Next = next
		pre = start
		end = pre
	}
	return dummy.Next
}