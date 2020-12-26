package leetcode242

import "sort"

// 两个单词如果包含相同的字母，次序不同，则称为字母易位词(anagram)
func isAnagram(s string, t string) bool {
	input1 := []byte(s)
	input2 := []byte(t)

	sort.SliceStable(input1, func(i, j int) bool {
		return input1[i] < input1[j]
	})

	sort.SliceStable(input2, func(i, j int) bool {
		return input2[i] < input2[j]
	})

	str1 := string(input1)
	str2 := string(input2)

	if str1 == str2 {
		return true
	}
	return false
	
}


// map
func isAnagram2(s string, t string) bool {
	arg1 := []byte(s)
	arg1Map := make(map[byte]int)
	for _,k := range arg1 {
		if v,ok := arg1Map[k]; ok {
			arg1Map[k] = v + 1
		}else {
			arg1Map[k] = 1
		}
	}

	arg2 := []byte(t)
	for _,b := range arg2 {
		if v,ok := arg1Map[b]; ok  {
			arg1Map[b] = v - 1
		}else {
			return false
		}
	}
	
	
	for _,v := range arg1Map { 
		if v != 0 {
			return false
		}
	}
	return true

}