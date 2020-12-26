package leetcode111 

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	ret := []int{}
	if root == nil {
		return 0
	}
	helper(root, 1, &ret)
	sort.SliceStable(ret, func(i, j int) bool {
		if ret[i] < ret[j] {
			return true
		} else {
			return false
		}
	})
	return ret[0]
}

// 把所有的情况都记录下来，然后排序求最小
func helper(node *TreeNode, count int, ret *[]int) {
	if node != nil && node.Left == nil && node.Right == nil {
		*ret = append(*ret, count)
		return
	}
	if node != nil {
		helper(node.Left, count+1, ret)
		helper(node.Right, count+1, ret)
	}
}



// 每次在求的过程中就开始比较大小 
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var miniCount *int  = new(int)
	*miniCount = math.MaxInt64
	helper2(root, 1, miniCount)
	return *miniCount
}

func helper2(node *TreeNode, count int, minCount *int) {
	if node != nil && node.Left == nil && node.Right == nil {
		if count < *minCount{
			*minCount = count
		}
		return
	}
	if node != nil {
		helper2(node.Left, count+1, minCount)
		helper2(node.Right, count+1, minCount)
	}
}


// solution3: 广度优先算法 
// 最先扫描到的叶子节点一定最小 BFS Beadth First Search  
func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	count := []int{}
	cur := root
	queue = append(queue, cur)
	count = append(count, 1)

	for i:=0; i < len(queue); i ++ { // 这步相当于给所有的元素标号，后面方便查找
		cur := queue[i]
		depth := count[i]

		if cur.Left == nil && cur.Right == nil {
			return depth
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
			count = append(count, depth + 1)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
			count = append(count, depth + 1)
		}
	}
	return 0
}