package leetcode

/*
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

// DP 方程：
// if S1[-1] != S2[-1] : LCS[s1,s2] =  max(LCS[s1-1,s2],LCS[s1,s2-1])
// if S1[-1] == S2[-1] : LCS[s1,s2] = LCS[s1-1,s2-1] + 1

// dp-table solution
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// longestCommonSubsequence2
func longestCommonSubsequence2(text1 string, text2 string) int {
	type node struct {
		i int
		j int
	}
	memo := map[node]int{}

	var dp func(int, int) int
	dp = func(i int, j int) int {
		if i == -1 || j == -1 {
			return 0
		}

		newNode := node{i: i, j: j}
		if n, ok := memo[newNode]; ok {
			return n
		}

		if text1[i] == text2[j] {
			return dp(i-1, j-1) + 1
		} else {
			maxNum := max(dp(i-1, j)+1, dp(i, j-1)+1)
			memo[newNode] = maxNum
			return maxNum
		}
	}

	return dp(len(text1)-1, len(text2)-1)
}
