package leetcode283

// https://leetcode-cn.com/problems/move-zeroes/

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[slow] != 0 {
			slow++
		} else {
			if nums[fast] != 0 {
				nums[fast], nums[slow] = nums[slow], nums[fast]
				slow++
			}
		}
		fast++
	}
}

// 暴力解法

func moveZeroes2(nums []int) {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i; j+1 < len(nums); j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			i--
			zeroCount++
		}
		if i == len(nums)-zeroCount-1 {
			return
		}

	}
}
