package leetcode

// 从前往后双指针，需要在内部重新开辟个空间
func merge(nums1 []int, m int, nums2 []int, n int) {
	ret := make([]int, 0, m+n) // notice: size = 0 capacity = n + n
	i, j := 0, 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			ret = append(ret, nums1[i])
			i++
		} else {
			ret = append(ret, nums2[j])
			j++
		}
	}

	if i == m {
		for ; j < n; j++ {
			ret = append(ret, nums2[j])
		}
	} else {
		for ; i < m; i++ {
			ret = append(ret, nums1[i])
		}
	}

	for i := 0; i < m+n; i++ {
		nums1[i] = ret[i]
	}
}

// 从后往前，不需要新开辟空间

func merge2(nums1 []int, m int, nums2 []int, n int) {
    p := m - 1 
    q := n - 1
    index := m + n - 1 
    
    for p >= 0 && q >=0 {
        if nums1[p] > nums2[q]{
            nums1[index] = nums1[p]
            p -- 
        }else {
            nums1[index] = nums2[q]
            q --
        }
        index -- 
    }

    for p<0 && q >= 0{
        nums1[index] = nums2[q]
        index -- 
        q -- 
    }

}
