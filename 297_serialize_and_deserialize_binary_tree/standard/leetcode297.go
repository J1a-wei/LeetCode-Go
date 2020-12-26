package leetcode297

import (
	"strconv"
	"strings"
)

/*
	会勾起大学学到的一些记忆，例如我们有了前序遍历和中序遍历，我们才能确定一棵树
	但是上述的这句话都是建立在没有写null的情况下，如果可以写null，我们也可以通过再走一遍前序遍历的顺序，实现反序列化

	本题先从pre DFS 开始尝试，后期可以把 前中后，层次的递归，非递归都练习一遍
*/

type TreeNode struct{
	Val int 
	Left *TreeNode
	Right *TreeNode
}


type Codec struct {
 l []string   
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    return rserialize(root,"")
}

// recursive serialize function 
func rserialize(root *TreeNode,str string)string {
	if root == nil {
		str += "null,"
	}else {
		str += strconv.Itoa(root.Val) + ","
		str = rserialize(root.Left,str)
		str = rserialize(root.Right,str)
	}
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
	l := strings.Split(data,",")
	for i:=0 ; i < len(l) ; i++ {
		if l[i] != "" {
			this.l = append(this.l, l[i])
		}
	}
	return this.rdeserialize()
}

/*
	1. 最开始的根，是可以返回回去的，除了最开始的根其他节点，都有节点去接住了，所以不需要担心递归无法返回第一次的结果 

	2. 队列如果是个nil，或者说队列长度是1，那么l[1:] 会不会抛异常？ 

*/
func (this *Codec) rdeserialize() *TreeNode{
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val: val}
	this.l = this.l[1:] 
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}