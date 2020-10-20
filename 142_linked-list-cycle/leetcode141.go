package main 

type ListNode struct{
	Val int 
	Next *ListNode
}
/*



	判断环原理: 

	定义 
		x1 令牌环head到环的第一个节点的步数，即我们要求的那个节点
		x2 从环的第一个点到相遇的那个点，需要x2步
		x3 从环中相遇点回到环的第一个点需要x3步 

	
	fast、slow会相遇，说明走的时间是一定的
	路程有如下关系
	fast : t = (x1 + x2 + x3 + x2) / 2
	slow:  t = (x1 + x2) / 1

	得到 x1 = x3 
	
	相遇后，我们让快指针指向头节点，然后步频为1的走，下次相遇即是x1的位置  


*/


func detectCycle(head *ListNode) *ListNode {
    fast,slow := head,head
	if head == nil{
		return nil
	}
	for {
		if fast.Next ==nil || fast.Next.Next == nil {
			return nil
		}
		
		fast = fast.Next.Next
		slow = slow.Next
        if fast == slow {
            break
        }
	}
    fast = head 
    for fast != slow {
        fast = fast.Next 
        slow = slow.Next
    }
    return fast

}