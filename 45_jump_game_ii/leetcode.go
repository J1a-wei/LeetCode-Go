package leetcode

func jump(nums []int) int {
	end := 0
	maxPosition := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			step++
		}

	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
