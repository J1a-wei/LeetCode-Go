package leetcode206

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 --> 2 --> 3 --> 4 --> 5 
// recusive 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head	
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil 
	return cur
}

// iterator 快慢指针 
func reverseList2(head *ListNode) *ListNode{
	var pre *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
