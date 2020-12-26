package leetcode

// https://leetcode-cn.com/problems/plus-one/submissions/
func plusOne(digits []int) (res []int) {
	index := len(digits) - 1
	for ; index >= 0; index-- {
		if digits[index] != 9 {
			break
		}
	}

	if index == -1 {
		res = append(res, 1)
		newSlice := make([]int, len(digits))
		res = append(res, newSlice...)
	} else {
		for i := len(digits) - 1; i > index; i-- {
			digits[i] = 0
		}
		digits[index]++
		res = digits
	}
	return
}
