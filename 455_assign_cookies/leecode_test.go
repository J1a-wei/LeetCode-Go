package leetcode

import "testing"

func TestLeetcode(t *testing.T) {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}

	result := findContentChildren(g, s)
	t.Log(result)
}
