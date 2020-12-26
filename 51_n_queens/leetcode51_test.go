package leetcode51

import "testing"

func TestLeetCode51(t *testing.T) {
	res := solveNQueens(9)
	t.Log(res)
	t.Log(len(res))
}
