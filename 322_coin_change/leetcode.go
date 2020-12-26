package leetcode

import (
	"sort"
)

// BFS
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	queue := []int{amount}

	visited := make([]bool, amount+1)
	visited[amount] = true
	step := 1
	sort.Ints(coins) // sort 方便剪枝
	for len(queue) != 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			head := queue[i]

			for _, v := range coins {
				next := head - v
				if next == 0 {
					return step
				}
				if next < 0 {
					break
				}

				if !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
			}

		}
		queue = queue[qLen:]
		step++
	}
	return -1
}

// DP

func coinChange2(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := 1; i < max; i++ {
		dp[i] = max
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
