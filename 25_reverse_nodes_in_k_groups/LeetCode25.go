package main


type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

// sub-recursive
func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}

// iterator
func reverseKGroup2(head *ListNode, k int) *ListNode{
	dummyHead := &ListNode{0, head}

	pre := dummyHead
	end := dummyHead


	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		start := pre.Next
		next := end.Next

		end.Next = nil 

		pre.Next = reverse2(start)
		start.Next = next
		
		
		pre = start 
		end = start 

	}
	return dummyHead.Next
}

// iterator
func reverse2(head *ListNode) *ListNode {
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
