package leetcode49

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string) 

	for _,s := range strs {
		bytes := []byte(s)
		sort.SliceStable(bytes,func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		if v,ok := m[string(bytes)]; ok {
			v = append(v, s)
			m[string(bytes)] = v  // map 应该是一个值类型的数据结构，按照java写法这里应该不需要再拷贝的 
		}else {
			slice := []string{s}
			m[string(bytes)] = slice
		}
	}
	ret := make([][]string,0,len(strs))
	for _,s := range m {
		ret = append(ret, s)
	}
	return ret

}