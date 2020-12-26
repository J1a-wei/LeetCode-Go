package leetcode236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	LCA问题，最近公共祖先

	这个问题里面有好多子问题要研究，是一个比较好的递归题

	1. 树的递归要在自己脑海中形成递归树
	2. 递归退出条件，这个应该好解决，前两个if
	3. 如何把找到的节点(p或者q) return回去，
		if left != nil {
			return right
		}
		return left
	4. 找到祖先后，又如何把祖先返回回去
		幸亏题意说过，本题所有节点是不重复的，以至于不需要特殊处理，祖先节点
		通过 if left !=nil && right != nil 确定， 通过第三条返回出去


*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

/*
	利用hash表存储父亲节点，利用父节点从P节点开始不断往上跳，并记录已经访问过的节点。
	再从q不断往上跳，如果碰到p的已经访问过的节点，那么就是答案

*/

func lowestCommonAncestorHash(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}

		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Left)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		if !visited[q.Val] {
			q = parent[q.Val]
		} else {
			return q
		}
	}
	return nil
}
