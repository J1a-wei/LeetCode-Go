package leetcode

import "math"

// https://leetcode-cn.com/problems/edit-distance/

// memo 备忘录解法
type Node struct {
	i int
	j int
}

func minDistance(word1 string, word2 string) int {
	i, j := len(word1), len(word2)
	memo := map[Node]int{}
	return dp(i-1, j-1, []byte(word1), []byte(word2), memo)
}

func dp(i int, j int, s1, s2 []byte, memo map[Node]int) int {
	// base case:
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	newNode := Node{i: i, j: j}
	if exist, ok := memo[newNode]; ok {
		return exist
	}

	if s1[i] == s2[j] {
		return dp(i-1, j-1, s1, s2, memo)
	} else {
		minNum := min(dp(i-1, j, s1, s2, memo)+1, dp(i, j-1, s1, s2, memo)+1, dp(i-1, j-1, s1, s2, memo)+1)
		memo[newNode] = minNum
		return minNum
	}
}

func min(nums ...int) int {
	minNum := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if minNum > nums[i] {
			minNum = nums[i]
		}
	}
	return minNum
}

// dp-table 解法
// that is awesome solution for dp
func minDistance2(word1 string, word2 string) int {
	s1, s2 := []byte(word1), []byte(word2)
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// base case
	// dp[0,0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				minNum := min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
				dp[i][j] = minNum
			}
		}
	}

	return dp[m][n]

}
