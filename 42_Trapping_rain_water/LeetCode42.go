package leetcode42

// 双指针
func trap(height []int) int {
	res,left,right,maxLeft,maxRight := 0,0,len(height) - 1,0,0
	for left <= right{
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			}else{
				res += maxLeft - height[left]
			}
			left ++
		}else{
			if height[right] > maxRight {
				maxRight = height[right]
			}else{
				res += maxRight  - height[right]
			}
			right --
		}
	}
	
    return res
}

// stack的应用 
func trap2(height []int) int {
	result,current := 0,0
	stack := make([]int,0,len(height))
	for current < len(height){
		for len(stack) !=0 && height[current] > height[stack[len(stack) - 1]] {
			left := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if len(stack) ==0 { // 这两个之间无法接水
				break
			}
			distance := current  - stack[len(stack) - 1] - 1
			h := min(height[current],height[stack[len(stack) - 1]]) - height[left]
			result += distance * h
		}
		stack = append(stack, current)
		current++
	}


	return result
}

func min(x,y int )int{
	if x < y {
		return x
	}
	return y
}