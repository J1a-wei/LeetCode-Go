package leetcode515

import "math"

/*
您需要在二叉树的每一行中找到最大的值。

示例：

输入:

          1
         / \
        3   2
       / \   \
      5   3   9

输出: [1, 3, 9]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	BFS
	在这个过程中，一直进行slice的截取，造成额外的开销。
*/
func largestValues(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)
		max := math.MinInt64
		for i := 0; i < size; i++ {
			current := queue[i]
			if current.Val > max {
				max = current.Val
			}
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		ret = append(ret, max)
	}
	return

}

// BFS beats 78% 较之前的好一些

func largestValues2(root *TreeNode) (ret []int) {
	if root == nil {
		return nil
	}
	var res = []int{}
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var length = len(queue)
		var max = -1 << 63
		for i := 0; i < length; i++ {
			if max < queue[i].Val {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, max)
		queue = queue[length:]
	}
	return res
}

//  还有一种写法，每次生成新的空间 Beat 77%
func largestValues3(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)
		max := -1 << 63
		newQueue := make([]*TreeNode, 0, 2*size)
		for i := 0; i < size; i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}

			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}
		ret = append(ret, max)
		queue = newQueue
	}
	return
}
