package leetcode

import "testing"

func TestLeetCode(t *testing.T) {
	bytes := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	count := numIslands(bytes)
	t.Log(count)
}
