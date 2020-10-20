package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var next *ListNode = head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

// Iterator
func swapPairs2(head *ListNode) *ListNode {
	pre := &ListNode{0, head}
	tmp := pre

	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end  // 注意这步，经常遗漏！   跟上层进行关联的节点 
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return pre.Next

}
