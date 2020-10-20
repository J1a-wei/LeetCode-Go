package leetcode26

func removeDuplicates(nums []int) int {
    if(len(nums) == 0){
        return 0
    }
    index := 1 
    for i:=1;i< len(nums) ; i ++ {
        j := i - 1 
        if(nums[i] != nums[j]){
            nums[index] = nums[i]
            index ++
        }
    }
    return index 
}