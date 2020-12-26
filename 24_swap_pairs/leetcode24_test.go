package leetcode

import "testing"

func TestLeetCode(t *testing.T) {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}

	ret := swapPairs3(head)

	t.Log(ret)
}
