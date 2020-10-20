package main 

type ListNode struct{
	Val int 
	Next *ListNode
}

// 快慢指针法，让其两个指针速度不一样，如果有环一定会相撞
func hasCycle(head *ListNode) bool {
	fast,slow := head,head
	if head == nil{
		return false
	}
	for {
		if fast.Next ==nil || fast.Next.Next == nil {
			return false
		}
		
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}

	}

}