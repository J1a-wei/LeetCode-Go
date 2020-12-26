package leetcode

// 运算符的优先级是个小坑
// 本题主要考察如何把二维数组转变成一维
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])

	left, right := 0, m*n-1

	for left <= right {
		mid_id := (left + right) >> 1
		mid := matrix[mid_id/n][mid_id%n]
		if mid == target {
			return true
		} else {
			if target > mid {
				left = mid_id + 1
			} else {
				right = mid_id - 1
			}

		}

	}
	return false
}
