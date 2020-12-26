package leetcode78

/*

给定一组不含重复元素的整数数组nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func subsets(nums []int) (result [][]int) {
	k := len(nums)
	tmp := []int{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == k {
			result = append(result, append([]int(nil), tmp...))
			return
		}
		// 选择当前元素
		tmp = append(tmp, nums[cur])
		dfs(cur + 1)

		tmp = tmp[:len(tmp)-1] // 不选当前元素
		dfs(cur + 1)
	}
	return
}

// 迭代算法
func subsets2(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ { // mask < 2的n次方 ；
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 == 1 { //  mask二进制的第i+1位（从右到左）是否为1
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}
