package leetcode297

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	slice := convertTreeNodeToSlice(root)
	if slice == nil {
		return "nil"
	} else {
		strSlie := []string{}
		for _, v := range slice {
			if v != nil {
				strSlie = append(strSlie, strconv.Itoa(*v))
			} else {
				strSlie = append(strSlie, "nil")
			}
		}

		res := ""
		for i, s := range strSlie {
			if i != len(strSlie)-1 {
				res = res + s + ","
			} else {
				res = res + s
			}
		}

		return res
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strSlice := strings.Split(data, ",")
	if strSlice == nil || len(strSlice) == 0 || strSlice[0] == "nil" {
		return nil
	}

	return deserializeHelper(strSlice, 0, nil)
}

func deserializeHelper(saveNodes []string, depth int, parents []*TreeNode) *TreeNode {
	iteralSize := 1 << depth
	var root *TreeNode
	newParents := []*TreeNode{}
	if depth == 0 {
		root = createNode(saveNodes[0])
		newParents = append(newParents, root)
		saveNodes = saveNodes[1:]
		deserializeHelper(saveNodes, depth+1, newParents)
		return root
	}

	if saveNodesAllNil(saveNodes) { // 终止条件，当所有的saveNodes均为nil
		return nil
	}

	for _, v := range parents {
		for i := 0; i < iteralSize; i = i + 2 {
			if v != nil && len(saveNodes) >= 2 {
				v.Left = createNode(saveNodes[0])
				v.Right = createNode(saveNodes[1])
				newParents = append(newParents, v.Left)
				newParents = append(newParents, v.Right)
				saveNodes = saveNodes[2:]
				break
			}
		}
	}

	return deserializeHelper(saveNodes, depth+1, newParents)
}

func createNode(s string) *TreeNode {
	if s == "nil" {
		return nil
	} else {
		x, _ := strconv.Atoi(s)

		return &TreeNode{
			x,
			nil,
			nil,
		}
	}
}

func saveNodesAllNil(nodes []string) bool {
	for _, test := range nodes {
		if test != "nil" {
			return false
		}
	}
	return true

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// 标号的过程，实际上是层序遍历 BFS
func convertTreeNodeToSlice(root *TreeNode) []*int {
	ret := []*int{}
	convertTreeNodeToSliceHelper([]*TreeNode{root}, &ret)
	return ret
}

// go 内置的int 和 string 都不可以为nil
func convertTreeNodeToSliceHelper(nodes []*TreeNode, ret *[]*int) {
	if len(nodes) == 0 {
		return
	}
	nextNodes := []*TreeNode{}
	for _, node := range nodes {
		if node != nil {
			var x = new(int)
			*x = node.Val
			*ret = append(*ret, x)

			nextNodes = append(nextNodes, node.Left)
			nextNodes = append(nextNodes, node.Right)
		} else {
			*ret = append(*ret, nil)

		}
	}
	convertTreeNodeToSliceHelper(nextNodes, ret)

}
