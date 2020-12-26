package leetcode589 

type Node struct {
	Val int 
	Children []*Node
}

func preorder(root *Node)[]int {
	stack := make([]*Node,0,100)
	ret := make([]int,0,100)

	if root == nil {
		return ret
	}
	stack= append(stack, root)

	for len(stack) != 0 {
		ns := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret, ns.Val)
		for i:=len(ns.Children) -1; i >=0; i -- {
			stack = append(stack, ns.Children[i])
		}
	}
	return ret


}