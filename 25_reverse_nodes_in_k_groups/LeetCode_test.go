package leetcode

import "testing"

func TestReverseFunc(t *testing.T) {
	three := &ListNode{3, nil}
	two := &ListNode{2, three}
	one := &ListNode{
		1,
		two,
	}

	ret := reverse2(one)
	t.Log(ret)

}
