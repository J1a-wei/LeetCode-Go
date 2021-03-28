package leetcode590

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ret := []int{}
	helper(root, &ret)
	return ret
}

func helper(root *Node, ret *[]int) {
	if root != nil {
		for _, v := range root.Children {
			helper(v, ret)
		}
		*ret = append(*ret, root.Val)
	}
}

// 先序遍历是中左右， 后序遍历是 左右中
// 那么我们调整一下先序遍历， 让其变成 中右左(在栈中，依次遍历子节点就能完成这个操作) ， 然后把结果数组翻转，输出的结果就是左右中了
func postOrder(root *Node) []int {
	ret := []int{}
	stack := make([]*Node, 0, 100)

	if root == nil {
		return ret
	}

	stack = append(stack, root)

	for len(stack) != 0 {
		ns := stack[len(stack)-1]
		ret = append(ret, ns.Val)
		stack = stack[:len(stack)-1]
		for _, n := range ns.Children {
			stack = append(stack, n)
		}
	}

	reverse(ret)
	return ret
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
