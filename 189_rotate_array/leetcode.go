package leetcode 



// 旋转数组 
func rotate(nums []int,k int){
	for i := 0; i < k; i++ {
		last := nums[len(nums) - 1]
		for j := len(nums) - 1; j- 1 >= 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}


// 最优解，每个元素只移动了一次，每个元素都是一步到位
func rotate2(nums []int,k int){
	k = k % len(nums)
	for start,count := 0,0; count < len(nums);start ++ {
		current := start 
		pre := nums[start]
		for  {
			next := (current + k) % len(nums)
			tmp := nums[next]
			nums[next] = pre
			pre = tmp 
			current = next 
			count ++
			if start == current {   // java 当中的do-while语句，在这里只能 for 死循环加上最后判断，才能实现至少执行一次的效果
				break
			}
		}
	}
}