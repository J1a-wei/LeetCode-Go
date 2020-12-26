package leetcode98

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/validate-binary-search-tree/
/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func isValidBST(root *TreeNode) bool {
	result := []int{}
	helper(root, &result)

	for i := 0; i+1 < len(result); i++ {
		if result[i+1] <= result[i] {  // 等于的话就不属于搜索树，有重复元素 
			return false
		}

	}
	return true
}

// 中序遍历改编
func helper(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	helper(node.Left, ret)
	*ret = append(*ret, node.Val)
	helper(node.Right, ret)
}


// 迭代中序遍历，去判断 
func isValidBST2(root *TreeNode) bool {
	stack := make([]*TreeNode,0,100)
	cur := root 
	lastRet := math.MinInt64
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}else {
			cur = stack[len(stack) - 1]
			if cur.Val > lastRet {
				lastRet = cur.Val
			}else {
				return false
			}
			stack = stack[ :len(stack) - 1]
			cur = cur.Right
		}
	}
	
	return true
}