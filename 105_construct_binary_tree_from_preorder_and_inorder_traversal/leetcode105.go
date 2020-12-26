package leetcode105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	通过前序遍历结果和中序遍历结果，确定一颗二叉树
	树中不存在重复元素
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	i := 0

	// 这部操作可以优化的，空间换时间操作  faster than 100% user
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)

	// 这部操作可以优化的，空间换时间操作
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}

	return helper(preorder, m, 0, len(m)-1, 0, len(m)-1)
}

func helper(preorder []int, inorder map[int]int, leftPreorder, rightPreorder, leftInorder, RightInorder int) *TreeNode {
	if leftPreorder > rightPreorder {
		return nil
	}

	root := &TreeNode{preorder[leftPreorder], nil, nil}

	i := inorder[preorder[leftPreorder]]

	subRootSize := i - leftInorder

	root.Left = helper(preorder, inorder, leftPreorder+1, leftPreorder+subRootSize, leftInorder, i-1)
	root.Right = helper(preorder, inorder, leftPreorder+subRootSize+1, rightPreorder, i+1, RightInorder)

	return root
}
