package leetcode104

import "sort"

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DSF 递归
func maxDepth(root *TreeNode) int {
	ret := []int{}
	if root == nil {
		return 0
	}

	helper(root, 1, &ret)

	sort.SliceStable(ret, func(i, j int) bool {
		if ret[i] > ret[j] {
			return true
		}
		return false
	})

	return ret[0]

}

func helper(node *TreeNode, count int, ret *[]int) {
	if node != nil &&node.Left == nil && node.Right == nil {
		*ret = append(*ret, count)
		return
	}
	if node != nil {
		helper(node.Left, count+1, ret)
		helper(node.Right, count+1, ret)
	}
}
