package leetcode226  

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
// https://leetcode-cn.com/problems/invert-binary-tree/description/
func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(node *TreeNode) {
	// termination  

	if node == nil {
		return 
	}

	// process on 
	node.Left,node.Right = node.Right,node.Left 

	// drill down 
	helper(node.Left)
	helper(node.Right)
}
