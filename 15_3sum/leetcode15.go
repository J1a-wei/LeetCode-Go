package leetcode15

import "sort"

// https://leetcode-cn.com/problems/3sum/

// a + b + c = 0
// 去重问题难解决，如果用java，可以用Set容器装，go语言需要自己处理去重
// func threeSum(nums []int) [][]int {
// 	sort.Slice(nums,func(i, j int) bool {
// 		if nums[i] <= nums[j] {
// 			return true
// 		}
// 		return false
// 	})
// 	result := make([][]int,0)

// 	for i:=0 ; i < len(nums) - 2 ;  i ++ {
// 		a := nums[i]

// 		for j,k := i + 1, len(nums) - 1; j < k;{
// 			b := nums[j]
// 			c := nums[k]
// 			if a + b + c == 0 {
// 				combol := []int{a,b,c}
// 				result = append(result, combol)
// 				break
// 			}else if a + b + c < 0 {
// 				j ++
// 			}else {
// 				k --
// 			}

// 		}
// 	}
// 	return result
// }

func threeSum(nums []int)[][]int{
	res := [][]int{}

	counter := map[int]int{}

	for _,value := range nums {
		counter[value] ++
	}
	uniques := []int{}
	for key := range  counter {
		uniques = append(uniques,key)
	}

	sort.Ints(uniques)

	for i := 0; i < len(uniques); i++ {
		if uniques[i] * 3 == 0 && counter[uniques[i]] >= 3 {
			res = append(res, []int{uniques[i],uniques[i],uniques[i]})
		}

		for j := i + 1; j < len(uniques); j++ {
			if uniques[i] * 2 + uniques[j] == 0 && counter[uniques[i]] > 1 {
				res = append(res, []int{uniques[i],uniques[i],uniques[j]})
			}

			if uniques[j] * 2 + uniques[i] == 0 && counter[uniques[j]] > 1 {
				res = append(res, []int{uniques[i],uniques[j],uniques[j]})
			}
			
			c := 0 - uniques[i] - uniques[j]
			if c > uniques[j] && counter[c] > 0 {
				res = append(res, []int{uniques[i],uniques[j],c})
			}

		}
	
	}

	return res
}



