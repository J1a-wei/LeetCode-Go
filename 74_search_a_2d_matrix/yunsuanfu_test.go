package leetcode

import "testing"

// 结论：
// +- 小于 位运算符的优先级  大于 */
func TestBitMove(t *testing.T) {
	res1 := 3 + 1>>1
	res2 := (3 + 1) >> 1
	res3 := 2 / 4 >> 1
	t.Log(res1, res2, res3)

}
