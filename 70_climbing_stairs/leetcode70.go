package leetcode70

// https://leetcode-cn.com/problems/climbing-stairs/

// 递归做法
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)

}

// 递归，带有缓存
func climbStairs3(n int) int {
	memo := make([]int, n+1)       // 3
	return climbStairsSub(n, memo) // 0, 2
}

func climbStairsSub(n int, memo []int) int {
	if n <= 1 {
		return n
	}

	if memo[n] == 0 {
		memo[n] = climbStairsSub(n-1, memo) + climbStairsSub(n-2, memo)
	}
	return memo[n]
}

// 迭代做法
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	result := make([]int, n+1, n+1)
	result[1] = 1
	result[2] = 2
	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}
