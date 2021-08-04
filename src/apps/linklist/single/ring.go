package single

import (
	"apps/linklist/models"
)

func HasRingByMap(head *models.ListNode) bool {
	if head == nil {
		return false
	}
	mp := make(map[*models.ListNode]struct{})
	for head != nil  {
		if _,ok := mp[head];ok {
			return true
		} else {
			mp[head] = struct{}{}

		}
		head = head.Next
	}
	return false
}

func HasRingByFastSlowPointer(head *models.ListNode) bool {

	return false
}



func max(x,y int) int{
	if x>y {
		return x
	}else {
		return y
	}

}