package leetcode283 



// https://leetcode-cn.com/problems/move-zeroes/

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int){
	if len(nums) ==0 {
		return
	}
	slow,fast := 0,0

	for fast < len(nums){
		if nums[slow] != 0 {
			slow ++
		}else{
			if nums[fast] != 0 {
				nums[fast],nums[slow] = nums[slow],nums[fast]
				slow ++
			}
		}
		fast ++
	}
}


// 左边找到一个等于0的，右边找一个不等于0的 进行交换，但无法保证相对位置有序 
func moveZeroesForce(nums []int){
	if len(nums) ==0 {
		return
	}
	// slow,fast := 0,0


}