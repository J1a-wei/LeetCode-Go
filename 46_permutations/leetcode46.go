package leetcode46

// https://leetcode-cn.com/problems/permutations/

/*

给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

排列讲究顺序的，不同顺序代表不同排列

回溯算法的O 是指数级别的
*/

func permute(nums []int) [][]int {
	length := len(nums)
	result := [][]int{}
	if length == 0 {
		return result
	}
	path := []int{}
	used := make([]bool, length, length) // 一开始默认false
	var dfs func(int)
	dfs = func(depth int) {
		if depth == length {
			// path 是公共变量，这里如果不去重新allocate一块内存，会导致数据错乱
			copyPath := make([]int, length)
			copy(copyPath, path)
			result = append(result, copyPath)
			return
		}
		for i := 0; i < length; i++ {
			if used[i] { // 已经用过
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs(depth + 1)

			path = path[:len(path)-1] // 这两部操作是回溯
			used[i] = false
		}
	}
	dfs(0)
	return result
}
