package leetcode

import "testing"

func TestLeetCode21(t *testing.T){
	input1 := &ListNode{
		2,
		nil,
	}
	input2 := &ListNode{
		1,
		nil,
	}
	
	tests := struct{
		name string
		arg1 *ListNode
		arg2 *ListNode
		want *ListNode
	}{
		"merge two sorted list",
		input1,
		input2,
		nil,
	}

	t.Run("test",func(t *testing.T) {
		ret := mergeTwoLists(tests.arg1,tests.arg2)
		for ret.Next != nil {
			t.Log(ret.Val)
		}
	})


}