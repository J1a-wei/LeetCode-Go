package leetcode169

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于⌊ n/2 ⌋的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例1:

输入: [3,2,3]
输出: 3
示例2:

输入: [2,2,1,1,1,2,2]
输出: 2


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/

// Map 方案
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

// Recursive,这个写法实质上就是把迭代改成了递归，而且仍然使用了map
func majorityElement2(nums []int) (ret int) {
	m := make(map[int]int)

	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			return
		}
		m[nums[i]]++
		if m[nums[i]] == len(nums)/2 {
			ret = nums[i]
		}

		dfs(i + 1)
	}
	dfs(0)
	return
}

/*
	分治法:
		众数，x >= len(nums)/2

	那么把nums从中间分开，众数至少会被分到一个part当中，并且在那个part仍然是众数
*/
func majorityElement3(nums []int) (ret int) {
	var backtrace func(int, int) int
	backtrace = func(lo int, hi int) int {
		if lo == hi {
			return nums[lo]
		}

		mid := (hi-lo)/2 + lo
		left := backtrace(lo, mid)
		right := backtrace(mid+1, hi)

		if left == right {
			return left
		}

		leftCount := countInRange(nums, left, lo, hi)
		rightCount := countInRange(nums, right, lo, hi)

		if leftCount > rightCount {
			return left
		}
		return right
	}

	return backtrace(0, len(nums)-1)
}

func countInRange(nums []int, num, left, right int) (count int) {
	for i := left; i <= right; i++ {
		if num == nums[i] {
			count++
		}
	}
	return
}
