package leetcode

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。

请找出其中最小的元素。



示例 1：

输入：nums = [3,4,5,1,2]
输出：1
示例 2：

输入：nums = [4,5,6,7,0,1,2]
输出：0
示例 3：

输入：nums = [1]
输出：1




来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := left + (right-left)>>1
		// 单调产生点
		if nums[mid] > nums[mid+1] {
			return nums[mid]
		}

		if mid-1 >= 0 && nums[mid-1] > nums[mid] {
			return nums[mid-1]
		}

		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		middle := (left + right) >> 1
		if nums[middle] < nums[right] { // middle 可能是最小值
			right = middle
		} else {
			// middle 一定不是最小值
			// 有点看不懂，即使middle不是最小值，为什么最小值一定在右边产生呢？
			// 因为右边出现了 nums[x + 1] < nums[x] 的情况,而这种打破单调的情况就是最小值产生点
			left = middle + 1
		}
	}
	return nums[left] // right也可以通过
}
