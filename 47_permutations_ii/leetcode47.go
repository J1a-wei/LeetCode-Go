package leetcode47

import (
	"sort"
)

// https://leetcode-cn.com/problems/permutations-ii/

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int) // backtrack 递归 走回头路
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] { // 这个剪枝条件有点难想出来，建议画图看下，为什么第二次选了跟上层一样，并且没被used的情况下会被重复
				continue
			}
			perm = append(perm, v)

			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
