package leetcode

import "math"

// 暴力遍历法
func maxSlidingWindowForce(nums []int, k int) []int {

	result := []int{}

	count := len(nums)

	for i := 0; i < count-k + 1; i++ {
		tmp := math.MinInt64
		for j := 0; j < k; j++ {
			if tmp < nums[i+j]{
				tmp = nums[i + j]
			}
		}
		result = append(result, tmp)
	}
	return result

}

// 队列解法 错误解法
func maxSlidingWindowQueue(nums []int, k int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}
	list := []int{}
	result := make([]int,len(nums) - k + 1)
	

	for i := 0; i < len(nums); i++ {
		for(len(list) != 0 && nums[list[len(list) - 1]] <= nums[i]){ // delete all elements that is smaller newNum
			// pop 
			list = list[:len(list) - 1]
		}
		list = append(list, i)

		if list[len(list) - 1] <= i -k { // queue full then delete one element
			list = list[1:]
		}

		if (i - k + 1 >=0) {
			result[i - k + 1] = nums[list[len(list) - 1]]
		}

	}
		return result

}

// 解法二 双端队列 
func maxSlidingWindows(nums []int,k int) []int{
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int,0,k)
	result := make([]int,0,len(nums) - k + 1)

	for i,v := range nums{
		// <=  实际上== 也可以 
		if i >= k  && window[0] <= i - k { // if the left-most index is out of window, remove it
			window = window[1:]
		}

		for len(window) > 0 && nums[window[len(window) - 1]] < v {
			window = window[:len(window) - 1]
		}

		window = append(window, i) // store the index of nums 

		if i >= k - 1 {
			result = append(result, nums[window[0]])
		}


	}
	return result
}