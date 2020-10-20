package leetcode84

import "math"

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
// 柱形图中最大的矩形

// 暴力解法 n^2 枚举宽
func largestRectangleAreaForce(heights []int) int {
	    result := 0 
    for i := 0; i < len(heights); i ++ {
        h := heights[i]
        l,r := i,i
        for ; l >=0 ; l -- {
            if heights[l] < heights[i]{
                break
            }
        }

        for;r < len(heights); r ++ {
            if heights[r] < heights[i]{
                break
            }
        }

        area :=  (r - l  - 1) * h
        result = max(area,result) 
    }
    return result

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

func largestRectangleAreaForceStackPrioriy(heights []int) int {
	n := len(heights)

	left , right := make([]int,n),make([]int ,n)

	for i := 0; i < n; i++ {  // 如果后面元素呈现递增势，相当于右挡板都是最后那个元素 n, 我们的最优解法只会判断一次，先把初始化，假设所有右挡板都是n
		right[i] = n
	}

	mono_stack := []int{}
	for i := 0 ; i < n; i ++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack) - 1]] >= heights[i]{
			right[mono_stack[len(mono_stack) - 1]] = i
			mono_stack = mono_stack[: len(mono_stack) - 1]
		}

		if len(mono_stack) == 0 {
			left[i] = -1
		}else{
			left[i] = mono_stack[len(mono_stack) - 1]
		}
		mono_stack = append(mono_stack, i)
	}

	result := 0

	for i := 0; i < n; i++ {
		result = max(result,heights[i] * (right[i] - left[i] - 1))
	}
	return result

}