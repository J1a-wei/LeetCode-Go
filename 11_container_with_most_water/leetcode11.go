package leetcode11

// https://leetcode-cn.com/problems/container-with-most-water/

// area distance * minHeight 
// best method
func maxArea(height []int) int {

	if len(height) < 2 {
		return 0
	}

	result,left,right := 0,0,len(height) - 1
	result = (right -  left) * min(height[left],height[right]) 
	for left < right {
		if height[left] < height[right]{
			left ++
		}else {
			right --

		}
		thisResult := (right -  left) * min(height[left],height[right]) 
		result = max(result,thisResult)
	}
	return result
}

func min(x,y int) int{
	if x < y {
		return x
	}
	return y
}

func max(x,y int) int{
	if x > y {
		return x
	}
	return y
}




// force 
func maxAreaForce(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}

	result := 0
	for i := 0; i < n ; i ++ {
		for j := 0; j <n;j ++ {
			if height[j] < height[i] || j == i {
				continue
			}
			dV := distance(j , i)
			result = max(result,dV * height[i])
		}
	}
	return result
}

func distance(x,y int)int {
    if x < y {
        return y - x
    }else{
        return x - y
    }
}