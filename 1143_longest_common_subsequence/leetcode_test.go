package leetcode

import "testing"

func TestLongestCommonSub(t *testing.T) {
	str1 := "hello world! "
	s := str1[0]
	t.Logf("%T", s)

	for i, v := range str1 {
		t.Logf("%T,%T,%v,%v", i, v, i, v)
	}

}
