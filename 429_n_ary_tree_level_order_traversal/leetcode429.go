package leetcode429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	queue := make([]*Node, 0, 50)
	result := [][]int{}

	if root == nil {
		return result
	}

	queue = append(queue, root)

	for len(queue) != 0 {
		layer := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			delNode := queue[0]
			queue = queue[1:]
			layer = append(layer, delNode.Val)
			// delNode不会为node，why？  因为Children如果为nil，那么Children... 不会有任何元素(连nil也没有)
			queue = append(queue, delNode.Children...)
		}
		result = append(result, layer)
	}

	return result
}

// 递归解法
func levelOrder2(root *Node) [][]int {

	result := [][]int{}
	if root == nil {
		return result
	}
	helper(root, 0, &result)
	return result
}

func helper(n *Node, level int, result *[][]int) {
	if len(*result) == level { // 正常情况下，result长度应该大于 level， 如果出现等于或者小于（实际上不会出现小于的情况) ，说明是第一次遍历到，需要新增
		layer := []int{}
		*result = append(*result, layer)
	}
	(*result)[level] = append((*result)[level], n.Val)
	for _, sub := range n.Children { // 终止条件其实在这里
		helper(sub, level+1, result)
	}
}

// 递归解法Closure
func levelOrder3(root *Node) (result [][]int) {
	if root == nil {
		return
	}

	var bfs func(node *Node, level int)

	bfs = func(node *Node, level int) {
		if len(result) == level {
			layer := []int{}
			result = append(result, layer)
		}

		result[level] = append(result[level], node.Val)

		for _, n := range node.Children {
			if n != nil {
				bfs(n, level+1)
			}
		}
	}
	bfs(root, 0)
	return
}
