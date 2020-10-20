package leetcode

import "sort"



func twoSum(nums []int, target int) []int {
	m := make(map[int]int,len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	sort.Ints(nums)
	j := len(nums) - 1
	i := 0
	for i < j{
		if target> nums[i] + nums[j]{
			i ++

		}else if target < nums[i] + nums[j]{
			j -- 

		}else {
			return []int{m[nums[i]],m[nums[j]]}
		}
	}
	return nil 

	/*
		这是一个错误的题解
		由于数组未必有序，无法使用夹逼原理

		想要通过map来记录原来的顺序，进行夹逼，
		但无法排除两个相等的元素，map会发生替换操作

	*/
}

// 最优解  O(n)
func twoSum2(nums []int,target int)[]int {

	m := make(map[int]int,len(nums))

	for i,v := range nums {
		anthoer := target - v
		if e,ok := m[anthoer];ok {
			return []int{e,i}
		}
		m[v] = i
	}
	return nil

}

// 暴力 O^2
func twoSum3(nums []int, target int) []int {
    for i,_ := range nums {
        for k,_ := range nums {
            if i == k {
                continue
            }
            if nums[i] + nums[k] == target {
                return []int{i,k}
            }
        }
    }
    return nil
}