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
		tmp.Next = end // 注意这步，经常遗漏！   跟上层进行关联的节点
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return pre.Next

}

func swapPairs3(head *ListNode) *ListNode {
	dummyHead := &ListNode{-1, head}

	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		start := cur.Next
		end := cur.Next.Next
		next := end.Next
		end.Next = nil
		cur.Next = reverse(start)
		start.Next = next
		cur = start
	}

	return dummyHead.Next
}

func reverse(node *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
