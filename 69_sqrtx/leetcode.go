package leetcode

/*
left  = 1
right = 8
mid = 4

4 * 4 > 8

right = 3
left = 1
mid = 2

2 * 2 < 8
left = 3
right = 3



*/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 0, x

	for left <= right {
		mid := left + (right-left)>>1

		squre := mid * mid

		if x == squre {
			return mid
		} else if squre > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right

}
