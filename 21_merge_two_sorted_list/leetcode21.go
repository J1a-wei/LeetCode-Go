package leetcode 

type ListNode struct{
	Val int 
	Next *ListNode
}

// 这题最好不要再原链表中操作，因为l1 l2的大小不太好确定，会莫名增加很多case


// 迭代 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	retPre := &ListNode{0,nil}
	current := retPre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val{
			current.Next = l1 
			l1 = l1.Next
		}else {
			current.Next = l2 
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 == nil {
		current.Next = l2
	}else{
		current.Next = l1
	}
	return retPre.Next
}

func mergeTwoListsRecursive(l1 *ListNode,l2 *ListNode) *ListNode{
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}else if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next,l2)
		return l1
	}else {
		l2.Next = mergeTwoListsRecursive(l1,l2.Next)
		return l2
	}

}