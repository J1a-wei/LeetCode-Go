package leetcode94

// 94:中序遍历
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/comments/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归写法
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}

	inorder(root, &ret)
	return ret
}

func inorder(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, ret)
	*ret = append(*ret, root.Val)
	inorder(root.Right, ret)
}

// 栈+多一个指针
// 指针的作用: 使其完成先把左边的压完，然后再去读数据，最后压右边的流程
func inorderTraversalStack(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 100)
	ret := []int{}
	cur := root

	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, cur.Val)
			cur = cur.Right // 注意这里不是入栈，好几次惯性写错
		}
	}

	return ret
}

func inorderTraversalStack2(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			// 这种写法会导致原来输入的input数据结构改变，其实很不推荐
			cur.Left = nil

		} else {
			result = append(result, cur.Val)
			stack = stack[:len(stack)-1]
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
	}
	return
}
