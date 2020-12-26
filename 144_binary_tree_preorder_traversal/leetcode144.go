package leetcode144

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/submissions/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 
func preorderTraversal(root *TreeNode)[]int{
	ret := []int{}

	preorder(root,&ret)
	return ret
}

func preorder(root *TreeNode,ret *[]int){
	if root != nil {
		*ret = append(*ret, root.Val)
		preorder(root.Left,ret)
		preorder(root.Right,ret)
	}
}



// 迭代 
func preOrderTraversalStack(root *TreeNode)[]int {
	stack := make([]*TreeNode,0,100)
	ret := []int{}

	if root == nil {
		return ret
	}

	stack = append(stack, root)

	for len(stack) != 0 {
		cur := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return ret
}

