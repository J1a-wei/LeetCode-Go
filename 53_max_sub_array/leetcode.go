package leetcode

import "math"

// https://leetcode-cn.com/problems/maximum-subarray/
// 最大子序和 = 当前元素自身最大，或者 包含之前最大
func maxSubArray(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	maxNum := math.MinInt64
	for i := 1; i < n; i++ {
		nums[i] = max(nums[i], nums[i]+nums[i-1])
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}
	// 第一个元素没有参与比较
	if maxNum < nums[0] {
		return nums[0]
	}
	return maxNum

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
