package leetcode102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

The same as: 429

*/

// BFS解法 beats 5%
// 由于slice机制问题，这种写法耗时较长
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		layer := []int{}
		for i := 0; i < size; i++ {
			delNode := queue[0]
			layer = append(layer, delNode.Val)
			queue = queue[1:]
			if delNode.Left != nil {
				queue = append(queue, delNode.Left)
			}
			if delNode.Right != nil {
				queue = append(queue, delNode.Right)
			}
		}

		res = append(res, layer)
	}
	return
}

// DFS解法
func levelOrderDFS(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	var dfs func(int, *TreeNode)
	dfs = func(level int, node *TreeNode) {
		if node == nil { // Termination
			return
		}
		if len(res) == level {
			newLayer := []int{}
			res = append(res, newLayer)
		}
		res[level] = append(res[level], node.Val)
		dfs(level+1, node.Left)
		dfs(level+1, node.Right)
	}
	dfs(0, root)
	return
}
