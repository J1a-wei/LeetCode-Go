package leetcode


// https://leetcode-cn.com/problems/plus-one/submissions/
func plusOne(digits []int) []int {
    tag := -1 
    for i := len(digits) - 1; i >=0; i -- {
        if digits[i] != 9 {
            tag = i 
            break 
        }
    }

    if tag == -1 {
        res := make([]int,len(digits) + 1,len(digits) + 1)
        res[0] = 1
        return res
    }else {
        digits[tag] = digits[tag] + 1
        for tag = tag + 1;tag < len(digits);tag++ {
            digits[tag] = 0
        }
        return digits
    }
}
