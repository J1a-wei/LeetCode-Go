package leetcode84

import "math"

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
// 柱形图中最大的矩形

// 暴力解法 n^2 枚举宽
func largestRectangleAreaForce(heights []int) int {
	result := 0.0
	for i := 0; i < len(heights); i++ {
		var minHeight float64 = math.MaxFloat64
		for j := i; j < len(heights); j++ { // 这步实际上是固定了宽度
			minHeight = math.Min(minHeight, float64(heights[j]))
			result = math.Max(result, minHeight*float64(j-i+1))
		}
	}

	return int(result)

}

// 暴力解法 n^2 枚举高

func largestRectangleAreaForce2(heights []int) int {
	result := 0.0
	for mid := 0; mid < len(heights); mid++ {

		height := heights[mid]

		left, right := mid, mid

		for left-1 >= 0 && heights[left-1] >= height { // 这种写法要记住，我们直接让变量等于我们想要比对的元素即可
			left--

		}

		for right+1 < len(heights) && heights[right+1] >= height {
			right++
		}

		result = math.Max(result, float64(height)*float64(right-left+1))

	}

	return int(result)

}

// 单调栈解法，最优解
func largestRectangleAreaForceStack(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	mono_stack := []int{}

	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1] // 弹栈
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}

	mono_stack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1] // 弹栈
		}
		if len(mono_stack) == 0 {
			right[i] = n
		} else {
			right[i] = mono_stack[len(mono_stack) - 1]
		}
		mono_stack = append(mono_stack, i)
	}

	result := 0

	for i := 0; i < n; i++ {
		result = max(result,(right[i] - left[i] - 1) * heights[i])
	}
	return result

}


func max(x,y int)int {
	if x > y {
		return x
	}
	return y
}