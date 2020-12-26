package leetcode297

import "testing"

func TestSerialize(t *testing.T){
	root := &TreeNode{0,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}
	c := Codec{}
	msg := c.serialize(root)
	t.Log(msg)
	root = c.deserialize(msg)
	t.Log(root)
}